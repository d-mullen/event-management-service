package zenkit

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"strings"

	"github.com/TV4/logrus-stackdriver-formatter"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	LogTenantField    = "zing.tnt"
	LogUserField      = "zing.usr"
	LogServiceField   = "zing.svc"
	LogRequestIdField = "z_request_id" // matches graphql's contextLogger field

	// LogTypeField identifies the type of log entries to be added by a logger.
	LogTypeField = "zing.log.type"
	// LogTypeAudit is the type log type of audit log entries.
	LogTypeAudit         = "zing.log.audit"
	LogTypeApplication   = "zing.log.application"
	DefaultAuditLogLevel = logrus.InfoLevel
)

type zenkitAuditLoggerMarker struct{}

type ctxAuditLogger struct {
	logger *logrus.Entry
	fields logrus.Fields
}

var (
	ctxAuditLoggerKey = &zenkitAuditLoggerMarker{}
)

func ContextLogger(ctx context.Context) *logrus.Entry {
	entry := ctxlogrus.Extract(ctx)
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return entry
	}
	requestId := meta[LogRequestIdField]
	if requestId != nil {
		entry = entry.WithField(LogRequestIdField, strings.Join(requestId, " "))
	}
	return entry
}

func WithLogEntryTypeAudit(entry *logrus.Entry) *logrus.Entry {
	entry.Logger.SetLevel(DefaultAuditLogLevel)
	return entry.WithField(LogTypeField, LogTypeAudit)
}

func WithGrpcLogLevel(entry *logrus.Entry) *logrus.Entry {
	cfgLevel := viper.GetString(GrpcLogLevelConfig)
	level, err := logrus.ParseLevel(cfgLevel)
	if err != nil {
		entry.WithFields(logrus.Fields{
			"level": cfgLevel,
		}).Error("unable to parse grpc log level config; defaulted to WARN")
		level = logrus.WarnLevel
	}
	entry.Logger.SetLevel(level)
	return entry
}

func Logger(name string, opts ...func(*logrus.Entry) *logrus.Entry) *logrus.Entry {
	InitConfig(name) // necessary for proper logger configuration
	log := logrus.New()
	if viper.GetBool(LogStackdriverConfig) {
		log.Formatter = stackdriver.NewFormatter(
			stackdriver.WithService(name),
			// TODO: stackdriver.WithVersion
		)
	}
	cfgLevel := viper.GetString(LogLevelConfig)
	level, err := logrus.ParseLevel(cfgLevel)
	if err != nil {
		log.WithFields(logrus.Fields{
			"level":         cfgLevel,
			LogServiceField: name,
			LogTypeField:    LogTypeApplication,
		}).Error("unable to parse log level config; defaulted to INFO")
		level = logrus.InfoLevel
	}
	log.SetLevel(level)
	entry := log.WithField(LogTypeField, LogTypeApplication)
	for _, option := range opts {
		entry = option(entry)
	}
	return entry.WithFields(logrus.Fields{
		LogServiceField: name,
	})
}

func WithAuditLogger(ctx context.Context, entry *logrus.Entry) context.Context {
	l := &ctxAuditLogger{
		logger: entry,
		fields: logrus.Fields{},
	}
	return context.WithValue(ctx, ctxAuditLoggerKey, l)
}

func ContextAuditLogger(ctx context.Context) *logrus.Entry {
	l, ok := ctx.Value(ctxAuditLoggerKey).(*ctxAuditLogger)
	if !ok || l == nil {
		return logrus.NewEntry(
			&logrus.Logger{
				Out:       ioutil.Discard,
				Formatter: new(logrus.TextFormatter),
				Hooks:     make(logrus.LevelHooks),
				Level:     logrus.PanicLevel,
			})
	}

	fields := logrus.Fields{}

	tags := grpc_ctxtags.Extract(ctx)
	for k, v := range tags.Values() {
		fields[k] = v
	}

	for k, v := range l.fields {
		fields[k] = v
	}

	entry := l.logger.WithFields(fields)

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return entry
	}
	requestId := meta[LogRequestIdField]
	if requestId != nil {
		entry = entry.WithField(LogRequestIdField, strings.Join(requestId, " "))
	}

	return entry
}

func AuditLogStreamServerInterceptor(entry *logrus.Entry) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		newCtx := WithAuditLogger(ss.Context(), entry)
		wrapped := grpc_middleware.WrapServerStream(ss)
		wrapped.WrappedContext = newCtx

		return handler(srv, wrapped)
	}
}

func AuditLogUnaryServerInterceptor(entry *logrus.Entry) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := WithAuditLogger(ctx, entry)

		resp, err = handler(newCtx, req)
		return
	}
}
