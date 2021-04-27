package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/zenoss/event-management-service/config"
	"github.com/zenoss/event-management-service/metrics"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
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

func init() {
	rand.Seed(time.Now().UnixNano())
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

	// run
	err := zenkit.RunGRPCServerWithEndpoint(ctx, ServiceName, func(svr *grpc.Server) error {

		if viper.GetBool(config.EventManagementEnabledConfig) {
			log.Debug("registering event management server")
			svc, err := NewEventManagementService(ctx)
			if err != nil {
				return err
			}
			proto.RegisterEventManagementServer(svr, svc)
		}

		return nil

	}, proto.RegisterEventManagementHandlerFromEndpoint())
	if err != nil {
		log.WithError(err).Fatal("error running gRPC server")
	}
}
