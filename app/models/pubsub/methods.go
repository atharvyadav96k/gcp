package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/atharvyadav96k/gcp/common"
	bus_error "github.com/atharvyadav96k/gcp/common/error"
)

// Publish sends a message to the specified Pub/Sub topic.
//
// Parameters:
//   - ctx: request context (controls timeout/cancellation)
//   - topicName: name of the Pub/Sub topic
//   - payload: data to be serialized and published
//
// Returns:
//   - error if client is not initialized, serialization fails,
//     or publish confirmation fails.
//
// NOTE:
// This implementation waits for publish confirmation using res.Get(ctx).
// For high-throughput systems, consider an async version.
func (s *Service) Publish(ctx context.Context, topicName string, payload any) error {
	if s == nil || s.Client == nil {
		return bus_error.ErrPubSubClientNotInitialized
	}

	topic := s.Client.Topic(topicName)
	if topic == nil {
		return bus_error.ErrInvalidTopic
	}

	bytes, err := common.ToJSON(payload)
	if err != nil {
		return err
	}

	res := topic.Publish(ctx, &pubsub.Message{
		Data: bytes,
	})

	_, err = res.Get(ctx)
	return err
}

// Close shuts down the Pub/Sub client and releases resources.
//
// It is safe to call multiple times. If the client is not initialized,
// it returns nil.
func (s *Service) Close() error {
	if s != nil && s.Client != nil {
		return s.Client.Close()
	}
	return nil
}
