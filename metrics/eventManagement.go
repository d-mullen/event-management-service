package metrics

import (
	"github.com/zenoss/zenkit/v5"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// Note: Add the new metric and view here. Add the view to Allviews in view.go
var (
	// MSetStatusTimeMs set status metrics time taken in milliseconds
	MSetStatusTimeMs = stats.Float64("set status/time_taken", "The time taken in millis for setting status", "ms")

	// SetStatusTimeMsView  set status times taken
	SetStatusTimeMsView = &view.View{
		Name:        "set status/searchTime",
		Measure:     MSetStatusTimeMs,
		Description: "The distribution of set status times taken",
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
	}
)
