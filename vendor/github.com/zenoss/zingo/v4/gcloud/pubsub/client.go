package pubsub

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"
)

// Client defines the interface for interacting with PubSub
type Client interface {
	Close() error
	Topic(string) Topic
	CreateTopic(context.Context, string) (Topic, error)
	Subscription(string) Subscription
	CreateSubscription(ctx context.Context, subscription string, topic Topic, ackDeadline time.Duration) (Subscription, error)
}

// NewClient creates a new default PubSub client for the given projectID
// It must be closed when it is no longer needed
func NewClient(ctx context.Context, projectID string) (Client, error) {
	return NewClientFromFactory(ctx, projectID, DefaultGoogleClientFactory)
}

func NewClientFromInternalFactory(ctx context.Context, projectID string, factory ClientFactory) (Client, error) {
	return factory(ctx, projectID)
}

func NewClientFromFactory(ctx context.Context, projectID string, factory GoogleClientFactory) (Client, error) {
	client, err := factory(ctx, projectID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new pubsub.Client")
	}
	return &defaultClient{client}, nil
}

// defaultPubSubClient is a wrapper for a pubsub.Client that implements our PubSubClient interface
type defaultClient struct {
	client *pubsub.Client
}

func (c *defaultClient) Close() error {
	return c.client.Close()
}

func (c *defaultClient) Topic(topicID string) Topic {
	return &defaultTopic{c.client.Topic(topicID)}
}

func (c *defaultClient) CreateTopic(ctx context.Context, topicID string) (Topic, error) {
	t, err := c.client.CreateTopic(ctx, topicID)
	return &defaultTopic{t}, err
}

func (c *defaultClient) Subscription(subID string) Subscription {
	return NewDefaultSubscription(c.client.Subscription(subID))
}

func (c *defaultClient) CreateSubscription(ctx context.Context, subID string, topic Topic, ackDeadline time.Duration) (Subscription, error) {
	sub, err := c.client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{
		Topic:       c.client.Topic(topic.ID()),
		AckDeadline: ackDeadline,
	})
	if err != nil {
		return nil, errors.Wrap(err, "pubsub client failed to create subscription")
	}
	return NewDefaultSubscription(sub), nil
}
