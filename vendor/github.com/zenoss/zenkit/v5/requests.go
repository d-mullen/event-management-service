package zenkit

import (
	"context"
	"sync/atomic"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	MConcurrentRequests = stats.Int64("grpc/concurrent_requests", "Concurrent requests", stats.UnitDimensionless)
	MIncomingRequests   = stats.Int64("grpc/incoming_requests", "Incoming requests", stats.UnitDimensionless)

	ConcurrentRequestsView = &view.View{
		Name:        "grpc/concurrent_requests",
		Description: "Number of concurrent requests",
		Measure:     MConcurrentRequests,
		Aggregation: view.LastValue(),
	}

	IncomingRequestsView = &view.View{
		Name:        "grpc/incoming_requests",
		Description: "Number of incoming requests",
		Measure:     MIncomingRequests,
		Aggregation: view.LastValue(),
	}
)

type RPCWrapper func(context.Context, func() error) error

func requestCounter(max int) func(context.Context, func() error) error {
	var (
		incoming int64
		counter  int64
		sem      *semaphore.Weighted
	)
	if max > 0 {
		sem = semaphore.NewWeighted(int64(max))
	}
	return func(ctx context.Context, handler func() error) error {
		stats.Record(ctx, MIncomingRequests.M(atomic.AddInt64(&incoming, 1)))
		defer func() { stats.Record(ctx, MIncomingRequests.M(atomic.AddInt64(&incoming, -1))) }()
		if max > 0 {
			if sem.TryAcquire(1) {
				defer sem.Release(1)
			} else {
				return status.Error(codes.Unavailable, "concurrent request limit reached")
			}
		}
		stats.Record(ctx, MConcurrentRequests.M(atomic.AddInt64(&counter, 1)))
		defer func() { stats.Record(ctx, MConcurrentRequests.M(atomic.AddInt64(&counter, -1))) }()
		return handler()
	}
}

func StreamServerInterceptorWithWrapper(wrapper RPCWrapper) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return wrapper(stream.Context(), func() error {
			return handler(srv, stream)
		})
	}
}

func UnaryServerInterceptorWithWrapper(wrapper RPCWrapper) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = wrapper(ctx, func() error {
			resp, err = handler(ctx, req)
			return err
		})
		return
	}
}

// ConcurrentRequestsMiddleware provides a unary and stream server interceptor that enforce a shared
// concurrent request policy.
func ConcurrentRequestsMiddleware(max int) (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor) {
	wrapper := requestCounter(max)
	return UnaryServerInterceptorWithWrapper(wrapper), StreamServerInterceptorWithWrapper(wrapper)
}
