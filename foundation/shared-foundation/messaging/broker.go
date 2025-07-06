package messaging

import "github.com/yourusername/shared-foundation/core"

// Broker interface for message queuing/pub-sub
type Broker interface {
	core.Service
	Publish(topic string, message []byte) error
	Subscribe(topic string, handler MessageHandler) error
	Unsubscribe(topic string) error
}

// MessageHandler processes messages from broker
type MessageHandler func(topic string, message []byte) error
