package pubsub

import (
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var (
	// DefaultGoogleClientFactory is the factory method for creating a Google PubSub client
	DefaultGoogleClientFactory GoogleClientFactory = pubsub.NewClient
	DefaultClientFactory       ClientFactory       = newClient

	// ErrConflictingSubscription is returned when creating a subscription and the
	// subscription ID already exists for another topic
	ErrConflictingSubscription = errors.New("Subscription with same ID already exists for different topic")

	// ErrMissingClient is returned when no client has been set
	ErrMissingClient = errors.New("No Client provided")
)

// GoogleClientFactory is a method that can give us a new Google pubsub.Client
type GoogleClientFactory func(context.Context, string, ...option.ClientOption) (*pubsub.Client, error)

type ClientFactory func(context.Context, string, ...option.ClientOption) (Client, error)

// Helper defines an interface for managing PubSub topics and subscriptions
type Helper interface {
	// Creates a PubSub Topic unless it already exists
	GetOrCreateTopic(context.Context, string) (Topic, error)

	// Creates a PubSub Subscription unless it already exists
	GetOrCreateSubscription(ctx context.Context, subscription string, topic Topic, ackDeadline time.Duration) (Subscription, error)
}

// NewHelper creates a new default PubSub helper that will use the PubSub Client from the context
func NewHelper(client Client) Helper {
	return &defaultHelper{client}
}

func newClient(ctx context.Context, projectID string, opts ...option.ClientOption) (Client, error) {
	client, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	return &defaultClient{client}, nil
}

type defaultHelper struct {
	client Client
}

func (h *defaultHelper) GetOrCreateTopic(ctx context.Context, topicName string) (Topic, error) {
	topic := h.client.Topic(topicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if topic exists")
	}

	if exists {
		logrus.WithField("topic", topicName).Info("Topic already exists")
	} else {
		logrus.WithField("topic", topicName).Info("Creating new topic")
		topic, err = h.client.CreateTopic(ctx, topicName)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create topic")
		}
	}
	return topic, nil
}

func (h *defaultHelper) GetOrCreateSubscription(ctx context.Context, subscription string, topic Topic, ackDeadline time.Duration) (Subscription, error) {
	logger := logrus.WithFields(logrus.Fields{
		"subscription": subscription,
		"topic":        topic.ID(),
	})

	// Check to see if the subscription already exists
	var sub Subscription
	sub = h.client.Subscription(subscription)
	exists, err := sub.Exists(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if subscription exists")
	}

	if exists {
		logger.Info("Subscription with that ID already exists")

		// Make sure the subscription is for the requested topic
		subtopic, err := sub.GetTopic(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to retrieve topic from existing subscription")
		}

		if subtopic.ID() != topic.ID() {
			return nil, errors.WithStack(ErrConflictingSubscription)
		}

	} else {
		logger.Info("Creating new subscription")

		// Ensure topic exists
		exists, err = topic.Exists(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to check that topic exists")
		}

		if !exists {
			topic, err = h.client.CreateTopic(ctx, topic.ID())
			if err != nil {
				return nil, errors.Wrap(err, "failed to create topic")
			}
		}

		// Create the new subscription
		sub, err = h.client.CreateSubscription(ctx, subscription, topic, ackDeadline)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create new subscription")
		}
	}
	return sub, nil
}
