package main

import (
	"context"
	"github.com/sirupsen/logrus"
	eventContextMongo "github.com/zenoss/event-management-service/pkg/adapters/datasources/eventcontext/mongo"
	eventTSGrpc "github.com/zenoss/event-management-service/pkg/adapters/datasources/eventts"
	eventQueryGrpc "github.com/zenoss/event-management-service/pkg/adapters/server/grpc"
	"math/rand"
	"time"

	"github.com/zenoss/event-management-service/config"
	"github.com/zenoss/event-management-service/metrics"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
	"github.com/zenoss/event-management-service/pkg/adapters/scopes/yamr"
	"github.com/zenoss/event-management-service/pkg/application/event"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/spf13/viper"
	yamrPb "github.com/zenoss/zing-proto/v11/go/cloud/yamr"
	"google.golang.org/grpc"

	. "github.com/zenoss/event-management-service/service"

	"github.com/zenoss/zenkit/v5"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_management"
	eventQueryProto "github.com/zenoss/zing-proto/v11/go/cloud/eventquery"
	"github.com/zenoss/zing-proto/v11/go/cloud/eventts"
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

	if err := zenkit.WaitUntilEnvoyReady(log); err != nil {
		log.WithError(err).Fatal("waiting for envoy failed")
	}

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

		if viper.GetBool(config.EventQueryEnabled) {
			cfg := MongoConfigFromEnv(nil)
			db, err := mongodb.NewMongoDatabaseConnection(ctx, cfg)
			if err != nil {
				log.WithField(logrus.ErrorKey, err).Error("failed to connect to mongodb")
				return err
			}
			eventContextAdapter, err := eventContextMongo.NewAdapter(ctx, cfg, db)
			if err != nil {
				log.WithField(logrus.ErrorKey, err).Error("failed to connect to event-ts-svc")
				return err
			}
			conn, err := zenkit.NewClientConnWithRetry(ctx, "event-ts-svc", zenkit.DefaultRetryOpts())
			if err != nil {
				log.Errorf("failed to get connection event-ts-svc: %q", err)
				return err
			}
			eventTSClient := eventts.NewEventTSServiceClient(conn)
			eventTSAdapter := eventTSGrpc.NewAdapter(eventTSClient)
			yamrAddress := "yamr-query-public" // TODO get this from config
			yamrConn, err := zenkit.NewClientConnWithRetry(ctx, yamrAddress, zenkit.DefaultRetryOpts())
			if err != nil {
				log.
					WithField(logrus.ErrorKey, err).
					Errorf("failed to get connection to %s: %q", yamrAddress, err)
				return err
			}
			yamrQueryClient := yamrPb.NewYamrQueryClient(yamrConn)
			entityScopeAdapter := yamr.NewAdapter(yamrQueryClient)
			eventApp := event.NewService(
				eventContextAdapter,
				eventTSAdapter,
				entityScopeAdapter)
			eventsQuerySvc := eventQueryGrpc.NewEventQueryService(eventApp)
			log.Debug("registering event query service")
			eventQueryProto.RegisterEventQueryServer(svr, eventsQuerySvc)
		}

		return nil

	}, proto.RegisterEventManagementHandlerFromEndpoint)
	if err != nil {
		log.WithError(err).Fatal("error running gRPC server")
	}
}
