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
	// MSearchTimeMs Search metrics time taken in milliseconds
	MSearchTimeMs = stats.Float64("search/time_taken", "The time taken in millis for store search", "ms")
	// MSearchCount Search metrics count of events read in.
	MSearchCount = stats.Int64("search/count", "The count of events returned by store search", "By")
	// MGetBulkEventsTimeMs Search metrics time taken in milliseconds
	MGetBulkEventsTimeMs = stats.Float64("search/time_taken", "The time taken in millis for store bulk events", "ms")
	// MGetBulkEventsCount Search metrics count of events read in.
	MGetBulkEventsCount = stats.Int64("search/count", "The count of events returned by store bulk events", "By")
	// MGetActiveEventsTimeMs Search metrics time taken in milliseconds
	MGetActiveEventsTimeMs = stats.Float64("search/time_taken", "The time taken in millis for store active events by entity", "ms")
	// MGetAllActiveEventsTimeMs Search metrics time taken in milliseconds
	MGetAllActiveEventsTimeMs = stats.Float64("search/time_taken", "The time taken in millis for store all active events", "ms")

	// SearchTimeMsView  search times taken
	SearchTimeMsView = &view.View{
		Name:        "store/searchTime",
		Measure:     MSearchTimeMs,
		Description: "The distribution of search times taken",
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
	}

	// SearchCountView search average count of events returned
	SearchCountView = &view.View{
		Name:        "store/searchCount",
		Measure:     MSearchCount,
		Description: "The number of events returned by searches",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
		Aggregation: view.Sum(),
	}

	// MGetBulkEventsTimeMsView  GEtBulkEvents times taken
	MGetBulkEventsTimeMsView = &view.View{
		Name:        "store/searchTime",
		Measure:     MGetBulkEventsTimeMs,
		Description: "The distribution of search times taken",
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
	}

	// MGetBulkEventsCountView GetBulkEvents average count of events returned
	MGetBulkEventsCountView = &view.View{
		Name:        "store/GetBulkEventsCount",
		Measure:     MGetBulkEventsCount,
		Description: "The number of events returned by GetBulkEvents",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
		Aggregation: view.Sum(),
	}

	// MGetActiveEventsTimeMsView  GEtBulkEvents times taken
	MGetActiveEventsTimeMsView = &view.View{
		Name:        "store/GetActiveEventsTime",
		Measure:     MGetActiveEventsTimeMs,
		Description: "The distribution of GetActiveEvents times taken",
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
	}

	// MGetAllActiveEventsTimeMsView  GEtBulkEvents times taken
	MGetAllActiveEventsTimeMsView = &view.View{
		Name:        "store/GetAllActiveEventsTime",
		Measure:     MGetAllActiveEventsTimeMs,
		Description: "The distribution of GetAllActiveEvents times taken",
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
	}
)
