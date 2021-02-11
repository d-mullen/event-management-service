package metrics

import (
	"go.opencensus.io/stats/view"
)

var AllViews = []*view.View{
	// Ingest
	IngestPubsubAckCount,
	IngestPubsubNackCount,
	IngestPubsubTimeSincePublish,
	IngestPubsubMsgProcessingTime,
	IngestPubSubNackTimeSincePublished,
	IngestPubSubNackTimeSinceReceived,
	IngestPubsubProcessEventLatency,
	IngestPubSubUnprocessedBacklog,

	// Store
	SearchTimeMsView,
	SearchCountView,
	MGetBulkEventsTimeMsView,
	MGetBulkEventsCountView,
	MGetActiveEventsTimeMsView,
	MGetAllActiveEventsTimeMsView,
}
