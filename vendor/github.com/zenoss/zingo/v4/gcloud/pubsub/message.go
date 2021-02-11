package pubsub

import (
	"cloud.google.com/go/pubsub"
)

type Message interface {
	Ack()
	Nack()
	GetMessage() *pubsub.Message
}

type defaultMessage struct {
	Message *pubsub.Message
}

func (m *defaultMessage) Ack() {
	m.Message.Ack()
}

func (m *defaultMessage) Nack() {
	m.Message.Nack()
}

func (m *defaultMessage) GetMessage() *pubsub.Message {
	return m.Message
}
