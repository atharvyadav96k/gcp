package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// InitPubSub initializes and returns a new Pub/Sub client.
//
// Parameters:
//   - ctx: context for cancellation and timeout
//   - projectId: Google Cloud project ID
//
// Returns:
//   - *pubsub.Client
//   - error if initialization fails
//
// NOTE:
// This function does not manage singleton state.
// Caller is responsible for reuse and closing the client.
func InitPubSub(ctx context.Context, projectId string) (*pubsub.Client, error) {
	return pubsub.NewClient(ctx, projectId)
}
