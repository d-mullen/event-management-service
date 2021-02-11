package metrics

import (
	"github.com/zenoss/zenkit/v5"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	MIngestPubsubAckCount               = stats.Int64("ingest/pubsub/ack_count", "Number of items acked from the ingest subscription", stats.UnitDimensionless)
	MIngestPubsubNackCount              = stats.Int64("ingest/pubsub/nack_count", "Number of items nacked from the ingest subscription", stats.UnitDimensionless)
	MIngestPubsubTimeSincePublish       = stats.Int64("ingest/pubsub/time_since_publish", "Time from when a message is published to pubsub until it is acked", stats.UnitMilliseconds)
	MIngestPubsubMsgProcessingTime      = stats.Int64("ingest/pubsub/process_time_since_received", "Time from when a message is received from the subscription until it is acked", stats.UnitMilliseconds)
	MIngestPubsubNackTimeSincePublished = stats.Int64("ingest/pubsub/nack_time_since_publish", "Time from when a message is published to pubsub until it is nacked", stats.UnitMilliseconds)
	MIngestPubsubNackTimeSinceReceived  = stats.Int64("ingest/pubsub/nack_time_since_received", "Time from when a message is received from the subscription until it is nacked", stats.UnitMilliseconds)
	MIngestPubsubProcessEventLatency    = stats.Int64("ingest/pubsub/process_event_latency", "Time spent during a processEvent call", stats.UnitMilliseconds)

	MIngestPubSubUnprocessedBacklog = stats.Int64("ingest/pubsub/unprocessed_backlog", "number of items in the unprocessed queue", stats.UnitDimensionless)
)

var (
	IngestPubsubAckCount = &view.View{
		Measure:     MIngestPubsubAckCount,
		Name:        "ingest/pubsub/ack_count",
		Description: "Total items acked from the ingest subscription",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, KeyWorker},
		Aggregation: view.Sum(),
	}

	IngestPubsubNackCount = &view.View{
		Measure:     MIngestPubSubUnprocessedBacklog,
		Name:        "ingest/pubsub/nack_count",
		Description: "Total items nacked from the ingest subscription",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, KeyWorker},
		Aggregation: view.Sum(),
	}

	IngestPubsubTimeSincePublish = &view.View{
		Name:        "ingest/pubsub/time_since_publish",
		Description: "Distribution of time from when an item is published to the ingest queue until it is acked",
		Measure:     MIngestPubsubTimeSincePublish,
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, KeyWorker},
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
	}

	IngestPubsubMsgProcessingTime = &view.View{
		Name:        "ingest/pubsub/time_since_received",
		Description: "Distribution of time from when an item is received from the subscription until it is acked",
		Measure:     MIngestPubsubMsgProcessingTime,
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, KeyWorker},
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
	}

	IngestPubsubProcessEventLatency = &view.View{
		Measure:     MIngestPubSubUnprocessedBacklog,
		Name:        "ingest/pubsub/process_event_latency",
		Description: "Distribution of time from when an item is processed for ingest",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, KeyWorker},
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
	}

	IngestPubSubUnprocessedBacklog = &view.View{
		Measure:     MIngestPubSubUnprocessedBacklog,
		Name:        "ingest/pubsub/unprocessed_backlog",
		Description: "Number of items in the unprocessed queue",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, KeyWorker},
		Aggregation: view.LastValue(),
	}

	IngestPubSubNackTimeSincePublished = &view.View{
		Measure:     MIngestPubsubNackTimeSincePublished,
		Name:        "ingest/pubsub/nack_time_since_publish",
		Description: "Time from when a message is published to pubsub until is nacked",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, KeyWorker},
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
	}

	IngestPubSubNackTimeSinceReceived = &view.View{
		Measure:     MIngestPubsubNackTimeSinceReceived,
		Name:        "ingest/pubsub/nack_time_since_received",
		Description: "Time from when a message is received from the subscription until is nacked",
		TagKeys:     []tag.Key{zenkit.TagKeyServiceLabel, KeyWorker},
		Aggregation: ocgrpc.DefaultMillisecondsDistribution,
	}
)
