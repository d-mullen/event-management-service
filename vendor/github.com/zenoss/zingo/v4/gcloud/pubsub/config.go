package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/zenoss/zenkit/v5"
)

type SubscriptionConfig struct {
	SubscriptionID string
	Config         pubsub.SubscriptionConfigToUpdate
}

type TopicConfig struct {
	TopicID       string
	Config        pubsub.TopicConfigToUpdate
	Subscriptions []SubscriptionConfig
}

func ZenkitInitializeTopics(ctx context.Context, confs ...TopicConfig) error {
	project := viper.GetString(zenkit.GCProjectIDConfig)
	return CreateTopicsFromConfig(ctx, project, confs...)
}

func CreateTopicsFromConfig(ctx context.Context, projectID string, confs ...TopicConfig) error {
	psClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	for _, topicConfig := range confs {
		topicID := topicConfig.TopicID
		requestedConfig := topicConfig.Config
		subscriptions := topicConfig.Subscriptions

		topic := psClient.Topic(topicID)
		exists, err := topic.Exists(ctx)
		if err != nil {
			return errors.Wrapf(err, "checking for topic %s existence", topicID)
		}
		var actualConfig pubsub.TopicConfig
		if !exists {
			topic, err = psClient.CreateTopic(ctx, topicID)
			if err != nil {
				return errors.Wrapf(err, "creating topic %s", topicID)
			}
		} else {
			actualConfig, err = topic.Config(ctx)
			if err != nil {
				return errors.Wrapf(err, "getting topic config for %s", topicID)
			}
		}
		if mapsDiffer(requestedConfig.Labels, actualConfig.Labels) {
			_, err = topic.Update(ctx, topicConfig.Config)
			if err != nil {
				return errors.Wrapf(err, "updating topic config for %s to %#v", topicID, requestedConfig)
			}
		}

		for _, subscriptionConfig := range subscriptions {
			subscriptionID := subscriptionConfig.SubscriptionID
			requestedConfig := subscriptionConfig.Config

			subscription := psClient.Subscription(subscriptionID)
			exists, err := subscription.Exists(ctx)
			if err != nil {
				return errors.Wrapf(err, "checking for subscription %s existence", subscriptionID)
			}
			if !exists {
				config := makeSubscriptionConfig(requestedConfig)
				config.Topic = topic
				_, err = psClient.CreateSubscription(ctx, subscriptionID, config)
				if err != nil {
					return errors.Wrapf(err, "creating subscription %s", subscriptionID)
				}
			} else {
				actualConfig, err := subscription.Config(ctx)
				if err != nil {
					return errors.Wrapf(err, "getting config for subscription %s", subscriptionID)
				}
				if actualConfig.Topic.ID() != topic.ID() {
					return errors.Errorf("Unexpected topic for subscription %s (expected %s, was %s)",
						subscriptionID, topic.ID(), actualConfig.Topic.ID())
				}
				changed, updateConfig := subscriptionConfigDiffs(actualConfig, makeSubscriptionConfig(requestedConfig))
				if changed {
					_, err = subscription.Update(ctx, updateConfig)
					if err == nil {
						return errors.Wrapf(err, "updating subscription config for %s to %#v",
							subscriptionID, updateConfig)
					}
				}
			}
		}
	}
	return nil
}

func makeSubscriptionConfig(src pubsub.SubscriptionConfigToUpdate) pubsub.SubscriptionConfig {
	retainAckedMessages, _ := src.RetainAckedMessages.(bool)
	var pushConfig pubsub.PushConfig
	if src.PushConfig != nil {
		pushConfig = *src.PushConfig
	}
	return pubsub.SubscriptionConfig{
		Topic:               nil,
		PushConfig:          pushConfig,
		AckDeadline:         src.AckDeadline,
		RetainAckedMessages: retainAckedMessages,
		RetentionDuration:   src.RetentionDuration,
		Labels:              src.Labels,
	}
}

func subscriptionConfigDiffs(s0, s1 pubsub.SubscriptionConfig) (bool, pubsub.SubscriptionConfigToUpdate) {
	changed := false
	out := pubsub.SubscriptionConfigToUpdate{}

	pcDiffer := s0.PushConfig.Endpoint != s1.PushConfig.Endpoint
	pcDiffer = pcDiffer || mapsDiffer(s0.PushConfig.Attributes, s1.PushConfig.Attributes)
	if pcDiffer {
		out.PushConfig = &s1.PushConfig
		changed = true
	}
	if s0.AckDeadline != s1.AckDeadline {
		out.AckDeadline = s1.AckDeadline
		changed = true
	}
	if s0.RetainAckedMessages != s1.RetainAckedMessages {
		out.RetainAckedMessages = s1.RetainAckedMessages
		changed = true
	}
	if s0.RetentionDuration != s1.RetentionDuration {
		out.RetentionDuration = s1.RetentionDuration
		changed = true
	}
	if mapsDiffer(s0.Labels, s1.Labels) {
		out.Labels = s1.Labels
		changed = true
	}
	return changed, out
}

func mapsDiffer(m0, m1 map[string]string) bool {
	if len(m0) != len(m1) {
		return true
	}
	for key, val := range m0 {
		if val != m1[key] {
			return true
		}
	}
	return false
}
