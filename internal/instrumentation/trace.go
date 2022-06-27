package instrumentation

import (
	"context"
	"encoding/json"
	"time"

	"github.com/spf13/viper"
	"github.com/zenoss/event-management-service/config"
	"github.com/zenoss/zenkit/v5"
	"go.opencensus.io/trace"
	"golang.org/x/time/rate"
)

var (
	limiterTraceSpan = rate.NewLimiter(rate.Every(50*time.Millisecond), 100)
)

func mustMarshal(v interface{}, errFunc func(error)) string {
	b, err := json.Marshal(v)
	if err != nil {
		if errFunc != nil {
			errFunc(err)
		}
		return ""
	}
	return string(b)
}

func StartSpan(ctx context.Context, name string) (context.Context, *trace.Span) {
	pZ := viper.GetFloat64(zenkit.TracingSampleRateConfig)
	traceSampler := trace.ProbabilitySampler(pZ)

	if viper.GetBool(config.TraceRateLimitedSamplingEnabled) && !limiterTraceSpan.Allow() {
		traceSampler = trace.NeverSample()
	}
	return trace.StartSpan(ctx, name, trace.WithSampler(traceSampler))
}

func AnnotateSpan(key, message string, span *trace.Span, attributes map[string]any, errHandler func(error)) {
	span.Annotate([]trace.Attribute{
		trace.StringAttribute(key, mustMarshal(attributes, errHandler)),
	}, "got request for active events")
}
