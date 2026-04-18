package pubsub

import (
	"sync"

	"cloud.google.com/go/pubsub"
)

// Service wraps the Google Cloud Pub/Sub client and manages its lifecycle.
//
// It ensures the client is initialized only once and reused across
// the application (singleton-style).
type Service struct {

	// Client is the underlying Pub/Sub client used for publishing
	// and subscribing to messages.
	Client *pubsub.Client

	// once ensures the client is initialized only once.
	once sync.Once
}
