package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
)

// Topic defines the interface for interacting with a pubsub topic
type Topic interface {
	Exists(context.Context) (bool, error)
	ID() string
	Delete(context.Context) error
	Subscriptions(context.Context) *pubsub.SubscriptionIterator
	Publish(ctx context.Context, message *pubsub.Message) *pubsub.PublishResult
	PublishByteArray(context.Context, []byte) *pubsub.PublishResult
	Stop()
}

type defaultTopic struct {
	topic *pubsub.Topic
}

func (t *defaultTopic) ID() string {
	return t.topic.ID()
}

func (t *defaultTopic) Exists(ctx context.Context) (bool, error) {
	return t.topic.Exists(ctx)
}

func (t *defaultTopic) Delete(ctx context.Context) error {
	return t.topic.Delete(ctx)
}

func (t *defaultTopic) Subscriptions(ctx context.Context) *pubsub.SubscriptionIterator {
	return t.topic.Subscriptions(ctx)
}

func (t *defaultTopic) Publish(ctx context.Context, message *pubsub.Message) *pubsub.PublishResult {
	return t.topic.Publish(ctx, message)
}

func (t *defaultTopic) PublishByteArray(ctx context.Context, messageData []byte) *pubsub.PublishResult {
	return t.Publish(ctx, &pubsub.Message{Data: messageData})
}

func (t *defaultTopic) Stop() {
	t.topic.Stop()
}
