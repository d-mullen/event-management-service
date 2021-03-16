package zenkit

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"regexp"
	"strings"

	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opencensus.io/tag"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceRegistrationFunc func(*grpc.Server) error

// RegisterEndpointFunc  function generated to register grpc gateway for transcoding
type RegisterEndpointFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

var unmarshalMessageRegex = regexp.MustCompile(`cannot unmarshal (?P<bad_type>\w+) into Go value of type (?P<good_type>\w+)`)
var missingDataMessageRegex = regexp.MustCompile(`looking for beginning of value`)
var codeSet = map[int]bool{
	int(codes.Unimplemented):   true,
	int(codes.InvalidArgument): true,
}

func RunGRPCServerWithHealth(ctx context.Context, name string, f ServiceRegistrationFunc) error {
	return runGRPCServer(ctx, name, f, true, nil)
}

func RunGRPCServer(ctx context.Context, name string, f ServiceRegistrationFunc) error {
	return runGRPCServer(ctx, name, f, false, nil)
}

func RunGRPCServerWithEndpoint(ctx context.Context, name string, f ServiceRegistrationFunc, ref RegisterEndpointFunc) error {
	return runGRPCServer(ctx, name, f, false, ref)
}

func WithTrapSIGINT(ctx context.Context, log *logrus.Entry) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:
			log.Info("SIGINT received")
			cancel()
		case <-ctx.Done():
		}
	}()
	return ctx
}

func runGRPCServer(ctx context.Context, name string, f ServiceRegistrationFunc, startHealth bool, endpointFunc RegisterEndpointFunc) error {
	// If useTLS is true, grpc will listen using tls cert.
	// If grpc based startHealth is true and useTLS is true, the grpc health will not use tls.

	InitConfig(name)

	// Add OpenCensus tag mutators to context.
	ctx, _ = tag.New(ctx, GetTagMutators()...)

	log := Logger(name)
	grpcLog := Logger(name, WithGrpcLogLevel)
	grpc_logrus.ReplaceGrpcLogger(grpcLog)

	auditLog := Logger(name, WithLogEntryTypeAudit)

	// Start listening for SIGINT
	ctx = WithTrapSIGINT(ctx, log)

	if startHealth {
		// Set up the health server
		health, err := RegisterHealthServer(ctx, log)
		if err != nil {
			return err
		}
		health.Serving()
		defer health.NotServing()
		defer health.Shutdown()
	}

	// Create a GRPC server
	grpcServer := NewGRPCServer(ctx, log, auditLog)
	if err := f(grpcServer); err != nil {
		return errors.Wrap(err, "unable to register service")
	}

	// Create a listener
	grpcAddr := viper.GetString(GRPCListenAddrConfig)
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return errors.Wrap(err, "unable to start listener")
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.WithError(err).Fatal("Error starting up grpc server")
		}
	}()

	log.WithField("grpc address", grpcAddr).Info("started grpc server")

	h2Addr := viper.GetString(HTTP2ListenAddrConfig)
	grpcTarget, err := url.Parse(fmt.Sprintf("http://localhost%v", grpcAddr))
	if err != nil {
		return errors.Wrap(err, "Error creating grpc target url")
	}

	skipJWTValidation := viper.GetBool(JWTSkip)
	if skipJWTValidation || viper.IsSet(JWKSURL) {

		h2cProxy := httputil.NewSingleHostReverseProxy(grpcTarget)
		h2cProxy.Director = proxyDirector(h2cProxy.Director)
		h2cProxy.Transport = &http2.Transport{
			// Allow http schema. This doesn't automatically disable TLS
			AllowHTTP: true,
			// Do disable TLS.
			DialTLS: func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(netw, addr)
			},
		}

		h2s := &http2.Server{}
		dialOpts := []grpc.DialOption{}
		dialOpts = append(dialOpts, grpc.WithInsecure())

		var endpointHandler http.Handler
		if endpointFunc == nil && !skipJWTValidation {
			endpointHandler, err = JWTHandler(ctx, h2cProxy)
			if err != nil {
				return err
			}
		} else {

			// grpc-gateway uses the field name in the proto by default. Change options for json pb marshalling to
			// use grpc standard that defaults to lower camel case for field names.
			// https://developers.google.com/protocol-buffers/docs/proto3#json_options
			// ZING-4277
			// Added HTTPProtoErrorHandler for ZING-6662 messsaging issue.
			gwmux := runtime.NewServeMux(
				runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false}),
				runtime.WithProtoErrorHandler(HTTPProtoErrorHandler))

			// handler that proxies to the grpc server, if json content
			// it goes through the grpc gateway for transcoding.
			// If not json it proxies directly to the grpce server
			endpointHandler = endpointHandlerFunc(h2cProxy, gwmux)

			if viper.IsSet(JWKSURL) && !skipJWTValidation {
				//if jwt validation parameters are set, wrap with
				//jwt validation
				endpointHandler, err = JWTHandler(ctx, endpointHandler)
				if err != nil {
					return err
				}
			}

			//grpc transcoder gateway set to proxy to ourselves on the grpc port
			err = endpointFunc(ctx, gwmux, grpcAddr, dialOpts)
			if err != nil {
				return errors.Wrap(err, "Could not register endpoint")
			}
		}

		// wrap to handle h2c requests
		endpointHandler = h2c.NewHandler(endpointHandler, h2s)

		srv := &http.Server{
			Addr:    h2Addr,
			Handler: endpointHandler,
		}
		// h2 listen with delegation to grpc or transcoder endpoint
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.WithError(err).Fatal("Error starting up proxy")
			}
		}()
		log.WithField("proxy address", h2Addr).Info("started endpoint proxy gateway")
	}

	// Wait for a cancel from upstream or a SIGINT
	<-ctx.Done()

	grpcServer.GracefulStop()
	log.Info("shut down server gracefully")

	return nil
}

func endpointHandlerFunc(grpcHandler http.Handler, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcHandler.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func proxyDirector(super func(*http.Request)) func(*http.Request) {
	return func(req *http.Request) {
		super(req)
		req.Host = req.URL.Host
	}
}

// HTTPProtoErrorHandler is a function that overrides standard Handler.
// Enhance the error message so that users can understand the API behaviour.
func HTTPProtoErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, request *http.Request, err error) {
	st, _ := status.FromError(err)
	message := st.Message()
	code := st.Code()

	// Early return if no err or error code is not one of the replaced ones below.
	if err == nil || !codeSet[int(code)] {
		runtime.DefaultHTTPProtoErrorHandler(ctx, mux, marshaler, w, request, err)
		return
	}

	var messageModified bool = false

	switch code {
	case codes.Unimplemented:
		// Clear up vague message: "Not Implemented"
		if message == "Not Implemented" {
			message = fmt.Sprintf("Action requested not supported by API. (%s)", message)
			messageModified = true
		}

	case codes.InvalidArgument:
		// Surgically replace vague or unclear messages and create a new status for runtime.DefaultHTTPProtoErrorHandler
		unmarshalMatch := unmarshalMessageRegex.FindStringSubmatch(message)
		missingDataMatch := missingDataMessageRegex.FindStringSubmatch(message)
		if unmarshalMatch != nil {
			// Replace vague message: "json: cannot unmarshal number into Go value of type string"
			result := make(map[string]string)
			for i, name := range unmarshalMessageRegex.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = unmarshalMatch[i]
				}
			}
			// Transform vague message: "cannot unmarshal number into Go value of type string"
			// to: "API doesn't accept a 'number' value for 'string' type property."
			message = fmt.Sprintf("API does not accept a '%s' value for '%s' type property.",
				result["bad_type"],
				result["good_type"])
			messageModified = true
		} else if missingDataMatch != nil {
			// Replace vague missing value message: "invalid character '}' looking for beginning of value"
			message = fmt.Sprintf("Missing value passed to the API. (%s)", message)
			messageModified = true
		}
	}

	if messageModified {
		s := status.New(code, message)
		err = s.Err()
	}
	runtime.DefaultHTTPProtoErrorHandler(ctx, mux, marshaler, w, request, err)
}
