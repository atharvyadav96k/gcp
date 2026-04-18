package app

import "context"

// PublishMessage publishes a message to the given Pub/Sub topic.
//
// Parameters:
//   - topic: name of the Pub/Sub topic
//   - payload: message data to be published (can be struct, map, etc.)
//
// NOTE:
// This function does not return an error and does not wait for confirmation.
// If you need delivery guarantees, use a version that handles PublishResult.
//
// Example:
//   app.PublishMessage("user-events", data)
func (a *App) PublishMessage(topic string, payload interface{}) {
	ctx := context.Background()

	if a.PubSub == nil {
		return
	}

	a.PubSub.Publish(ctx, topic, payload)
}
