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
	MSetStatusTimeMs = stats.Float64("event_management/set_status_time_taken", "The time taken in millis for setting status", "ms")
	// MSetStatusCount Search metrics count of events read in.
	MSetStatusCount = stats.Int64("event_management/set_status_count", "The count of event occurrences whose status were set", "By")
	// MAnnotateTimeMs annotate metrics time taken in milliseconds
	MAnnotateTimeMs = stats.Float64("event_management/annotate_time_taken", "The time taken in millis for annotating", "ms")
	// MAnnotateCount count of event annotaions added or edited.
	MAnnotateCount = stats.Int64("event_management/annotate_count", "The count of event occurrences whose status were set", "By")
	// MDeleteAnnotationCount count of event annotations deleted.
	MDeleteAnnotationCount = stats.Int64("event_management/delete_annotation_count", "The count of annotations deleted", "By")

	// SetStatusTimeMsView  set status times taken
	SetStatusTimeMsView = &view.View{
		Name:        "setStatusTime",
		Measure:     MSetStatusTimeMs,
		Description: "The distribution of set status times taken",
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
	}

	// SetStatusCountView search average count of events returned
	SetStatusCountView = &view.View{
		Name:        "setStatusCount",
		Measure:     MSetStatusCount,
		Description: "The number of event occurrences whose status was set",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
		Aggregation: view.Sum(),
	}

	// AnnotateTimeMsView  set status times taken
	AnnotateTimeMsView = &view.View{
		Name:        "annotateTime",
		Measure:     MAnnotateTimeMs,
		Description: "The distribution of annotate times taken",
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
	}

	// AnnotateCountView total count of annotations created or edited
	AnnotateCountView = &view.View{
		Name:        "annotateCount",
		Measure:     MAnnotateCount,
		Description: "The number of annotations added or edited",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
		Aggregation: view.Sum(),
	}

	// DeleteAnnotationCountView total count of annotations deleted
	DeleteAnnotationCountView = &view.View{
		Name:        "DeleteAnnotationCount",
		Measure:     MDeleteAnnotationCount,
		Description: "The number of annotations deleted",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, zenkit.KeyTenant, KeyWorker},
		Aggregation: view.Sum(),
	}
)
