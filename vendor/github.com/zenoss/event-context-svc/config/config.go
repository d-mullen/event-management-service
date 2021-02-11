package config

import (
	"time"

	cps "cloud.google.com/go/pubsub"

	"github.com/spf13/viper"
	"github.com/zenoss/zingo/v4/gcloud/pubsub"
)

const (
	DefaultIngestPubsubAckDeadline             = 2 * time.Minute
	DefaultIngestPubsubBatchSize               = 100
	DefaultIngestPubsubWorkers                 = 20
	DefaultIngestPubsubMaxWait                 = 200 * time.Millisecond
	DefaultIngestPubsubRetryTimeout            = 10 * time.Hour
	DefaultIngestPubsubGoroutines              = 2
	DefaultIngestPubsubMaxExtension            = 10 * time.Minute
	DefaultIngestPubsubMaxOutstandingMegabytes = 512
	DefaultIngestPubsubMaxOutstandingMessages  = 2 * DefaultIngestPubsubWorkers * DefaultIngestPubsubBatchSize
	DefaultIngestRequestTimeout                = 5 * time.Minute
	DefaultIngestPubsubTopic                   = "event-context-in"
	DefaultIngestPubsubSubscription            = "event-context-in_sub_event-context-ingest"
	DefaultPublishPubsubTopic                  = "event-context-out"
	DefaultIDCacheSize                         = 1000000
	DefaultStoreWorkers                        = 20

	// EventContextIngestEnabledConfig determines if the API endpoint for ingest is enabled for this server
	EventContextIngestEnabledConfig = "ingest.enabled"
	// EventContextIngestPubSubReaderEnabledConfig determines if the pub-sub reader subscription read loop is enabled for this server
	EventContextIngestPubSubReaderEnabledConfig = "ingest.pubsub.reader.enabled"
	// EventContextQueryEnabledConfig determines if the API endpoint for query is enabled for this server
	EventContextQueryEnabledConfig = "query.enabled"
	// EventContextIngestPubsubTopicConfig the pub-sub topic this server listens to ingest event messages
	EventContextIngestPubsubTopicConfig = "ingest.pubsub.topic"
	// EventContextIngestPubsubSubscriptionConfig the pub-sub subscriptions used to read event messages for ingest
	EventContextIngestPubsubSubscriptionConfig = "ingest.pubsub.subscription"
	// EventContextPublishPubsubTopicConfig the pub-sub topic to which this server publishes event messages
	EventContextPublishPubsubTopicConfig = "publish.pubsub.topic"
	// EventContextIngestPubsubAckDeadlineConfig the duration by which a pub-sub message must be acknowledged
	EventContextIngestPubsubAckDeadlineConfig = "ingest.pubsub.AckDeadline"
	// IngestLateDataThresholdConfig the duration in which a message can be ingested by this server
	IngestLateDataThresholdConfig = "ingest.late.data.threshold"
	// FirestoreProjectConfig the project in which the Cloud Firestore instance is found
	FirestoreProjectConfig = "firestore.project.id"
	// EventContextIngestIDCacheSizeConfig the size of the event ID cache
	EventContextIngestIDCacheSizeConfig = "ingest.id.cache.size"
	// EventContextStoreNumWorkersConfig configures the number of works in the context store
	EventContextStoreNumWorkersConfig = "store.num.workers"
)

// InitDefaults sets defaults values for this server's configuration
func init() {
	viper.SetDefault(EventContextIngestEnabledConfig, false)
	viper.SetDefault(EventContextQueryEnabledConfig, true)
	viper.SetDefault(EventContextIngestPubsubAckDeadlineConfig, 10*time.Second)
	viper.SetDefault(EventContextIngestPubsubTopicConfig, DefaultIngestPubsubTopic)
	viper.SetDefault(EventContextIngestPubsubSubscriptionConfig, DefaultIngestPubsubSubscription)
	viper.SetDefault(EventContextPublishPubsubTopicConfig, DefaultPublishPubsubTopic)
	viper.SetDefault(EventContextIngestIDCacheSizeConfig, DefaultIDCacheSize)
	viper.SetDefault(EventContextStoreNumWorkersConfig, DefaultStoreWorkers)
}

// GetIngestPubSubConfig returns a PubSubConfig
func GetIngestPubSubConfig() (pubsub.TopicConfig, bool) {
	topicName := viper.GetString(EventContextIngestPubsubTopicConfig)
	subscriptionName := viper.GetString(EventContextIngestPubsubSubscriptionConfig)
	if topicName == "" || subscriptionName == "" {
		return pubsub.TopicConfig{}, false
	}
	return pubsub.TopicConfig{
		TopicID: topicName,
		Subscriptions: []pubsub.SubscriptionConfig{
			{
				SubscriptionID: subscriptionName,
				Config: cps.SubscriptionConfigToUpdate{
					AckDeadline: viper.GetDuration(EventContextIngestPubsubAckDeadlineConfig),
				},
			},
		},
	}, true
}

// GetPublishPubSubConfig returns a PubSubConfig
func GetPublishPubSubConfig() (pubsub.TopicConfig, bool) {
	topicName := viper.GetString(EventContextPublishPubsubTopicConfig)
	if topicName == "" {
		return pubsub.TopicConfig{}, false
	}
	return pubsub.TopicConfig{
		TopicID: topicName,
	}, true
}
