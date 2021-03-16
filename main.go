package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/zenoss/event-management-service/config"
	"github.com/zenoss/event-management-service/metrics"

	"github.com/cenkalti/backoff"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	. "github.com/zenoss/event-management-service/service"

	"github.com/zenoss/zenkit/v5"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_management"
)

const (
	// ServiceName is the name if this microservice.
	ServiceName = "event-management-service"
)

var (
	initializeMode    bool
	initializeTimeout time.Duration
)

func init() {
	rand.Seed(time.Now().UnixNano())
	pflag.BoolVar(&initializeMode, "init", false, "Initialize external resources, then exit")
	pflag.DurationVar(&initializeTimeout, "init-timeout", 10*time.Second, "Timeout waiting for external resources")

	pflag.Parse()
}

func initialize(ctx context.Context) error {

	// Retry pubsub connection within the timeout
	boff := backoff.NewExponentialBackOff()
	boff.MaxElapsedTime = initializeTimeout
	retryFunc := func() error {
		//err := initializePubsub(ctx)
		//if err != nil {
		//	logrus.WithError(err).Warn("failed to initialize Pubsub, will retry until timeout")
		//}
		return nil //err
	}
	err := backoff.Retry(retryFunc, boff)
	if err != nil {
		return errors.Wrap(err, "failed to initialize Pubsub within the timeout")
	}
	return nil
}

func main() {
	zenkit.InitConfig(ServiceName)

	log := zenkit.Logger(ServiceName)
	ctx, cancel := context.WithCancel(ctxlogrus.ToContext(context.Background(), log))
	defer cancel()

	// Register metrics
	if err := zenkit.RegisterViews(ServiceName, metrics.AllViews...); err != nil {
		log.WithError(err).Fatal("Error registering views")
	}

	if initializeMode {
		log.Info("running in init mode")
		_ = initialize(ctx)
		return
	}

	err := zenkit.RunGRPCServer(ctx, ServiceName, func(svr *grpc.Server) error {

		if viper.GetBool(config.EventManagementEnabledConfig) {
			log.Debug("registering event management  server")
			svc, err := NewEventManagementService(ctx)
			if err != nil {
				return err
			}
			proto.RegisterEventManagementServer(svr, svc)
		}

		return nil

	})
	if err != nil {
		log.WithError(err).Fatal("error running gRPC server")
	}
}
