package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"
)

// Subscription defines the interface for interacting with a pubsub subscription
type Subscription interface {
	Exists(context.Context) (bool, error)
	GetTopic(context.Context) (Topic, error)
	Delete(context.Context) error
	Receive(context.Context, func(context.Context, Message)) error
	ReceiveWithSettings(context.Context, pubsub.ReceiveSettings, func(context.Context, Message)) error
}

// defaultSubscription is a wrapper for a pubsub.Subscription that implements our interface
type defaultSubscription struct {
	subscription *pubsub.Subscription
}

func NewDefaultSubscription(subscription *pubsub.Subscription) Subscription {
	return &defaultSubscription{subscription}
}

func (s *defaultSubscription) Exists(ctx context.Context) (bool, error) {
	return s.subscription.Exists(ctx)
}

func (s *defaultSubscription) GetTopic(ctx context.Context) (Topic, error) {
	config, err := s.subscription.Config(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get subscription config")
	}
	return &defaultTopic{config.Topic}, nil
}

func (s *defaultSubscription) Delete(ctx context.Context) error {
	return s.subscription.Delete(ctx)
}

func (s *defaultSubscription) Receive(ctx context.Context, f func(context.Context, Message)) error {
	innerFunc := func(ctx2 context.Context, message *pubsub.Message) {
		msg := &defaultMessage{Message: message}
		f(ctx2, msg)
	}
	return s.subscription.Receive(ctx, innerFunc)
}

func (s *defaultSubscription) ReceiveWithSettings(ctx context.Context, settings pubsub.ReceiveSettings, f func(context.Context, Message)) error {
	oldSettings := s.subscription.ReceiveSettings
	s.subscription.ReceiveSettings = settings
	defer func() {
		s.subscription.ReceiveSettings = oldSettings
	}()

	return s.Receive(ctx, f)
}
