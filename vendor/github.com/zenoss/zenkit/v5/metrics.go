package zenkit

import (
	"context"
	"strings"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/viper"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"google.golang.org/grpc"
	grpcstats "google.golang.org/grpc/stats"
)

var (
	TagKeyServiceLabel, _ = tag.NewKey("service_label")
)

// Register views, adds zenkit specific tags and prefix to view name if not already set.
// This modifies the passed in views. Prefix is generally set to the mircoservice name.
func RegisterViews(prefix string, views ...*view.View) error {
	AddZenKitTags(views...)
	if len(prefix) > 0 {
		viewPrefix := prefix + "/"
		for _, view := range views {
			if !strings.HasPrefix(view.Name, viewPrefix) {
				view.Name = viewPrefix + view.Name
			}
		}
	}
	return view.Register(views...)
}

func AddZenKitTags(views ...*view.View) {
	globalTagKeys := GetTagKeys()

	for _, view := range views {
		for _, globalTagKey := range globalTagKeys {
			found := false
			for _, viewTagKey := range view.TagKeys {
				if viewTagKey.Name() == globalTagKey.Name() {
					found = true
				}
			}
			if !found {
				view.TagKeys = append(view.TagKeys, globalTagKey)
			}
		}
	}
}

// GetTagKeys returns tag keys for all metrics.
func GetTagKeys() []tag.Key {
	keys := make([]tag.Key, 0, 4)

	// Add service_label if available.
	serviceLabel := viper.GetString(ServiceLabel)
	if serviceLabel != "" {
		keys = append(keys, TagKeyServiceLabel)
	}

	return keys
}

// GetTagMutators returns tag mutators for all metrics.
func GetTagMutators() []tag.Mutator {
	mutators := make([]tag.Mutator, 0, 4)

	// Add service_label if available.
	serviceLabel := viper.GetString(ServiceLabel)
	if serviceLabel != "" {
		mutators = append(mutators,
			tag.Upsert(TagKeyServiceLabel, serviceLabel))
	}

	return mutators
}

func MetricTagsStreamServerInterceptor() grpc.StreamServerInterceptor {
	mutators := GetTagMutators()
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapped := grpc_middleware.WrapServerStream(stream)
		ctx, _ := tag.New(stream.Context(), mutators...)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}
}

func MetricTagsUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	mutators := GetTagMutators()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, _ = tag.New(ctx, mutators...)
		return handler(ctx, req)
	}
}

type GrpcStatsHandler struct {
	Handler grpcstats.Handler
}

var _ grpcstats.Handler = &GrpcStatsHandler{}

func (s *GrpcStatsHandler) TagRPC(ctx context.Context, rti *grpcstats.RPCTagInfo) context.Context {
	ctx, _ = tag.New(ctx, GetTagMutators()...)
	return s.Handler.TagRPC(ctx, rti)
}

// HandleRPC processes the RPC stats.
func (s *GrpcStatsHandler) HandleRPC(ctx context.Context, rs grpcstats.RPCStats) {
	ctx, _ = tag.New(ctx, GetTagMutators()...)
	s.Handler.HandleRPC(ctx, rs)
}
func (s *GrpcStatsHandler) TagConn(ctx context.Context, cti *grpcstats.ConnTagInfo) context.Context {
	return s.Handler.TagConn(ctx, cti)
}

// HandleConn processes the Conn stats.
func (s *GrpcStatsHandler) HandleConn(ctx context.Context, cs grpcstats.ConnStats) {
	s.Handler.HandleConn(ctx, cs)
}

func SinceInMS(start time.Time) int64 {
	return int64(time.Since(start) / time.Millisecond)
}

func SinceInHours(start time.Time) int64 {
	return int64(time.Since(start) / time.Hour)
}

func SinceInMinutes(start time.Time) int64 {
	return int64(time.Since(start) / time.Minute)
}

func RecordMillisecondLatency(ctx context.Context, ms *stats.Int64Measure, since time.Time, mutator ...tag.Mutator) {
	ctx, _ = tag.New(ctx, mutator...)
	stats.Record(ctx, ms.M(SinceInMS(since)))
}

func RecordHourLatency(ctx context.Context, ms *stats.Int64Measure, since time.Time, mutator ...tag.Mutator) {
	ctx, _ = tag.New(ctx, mutator...)
	stats.Record(ctx, ms.M(SinceInHours(since)))
}

func RecordMinuteLatency(ctx context.Context, ms *stats.Int64Measure, since time.Time, mutator ...tag.Mutator) {
	ctx, _ = tag.New(ctx, mutator...)
	stats.Record(ctx, ms.M(SinceInMinutes(since)))
}
