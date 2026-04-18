package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

// InitFirestore creates and returns a new Firestore client for the given project.
//
// Parameters:
//   - projectId: Google Cloud project ID
//
// Returns:
//   - *firestore.Client: initialized Firestore client
//   - error: if client initialization fails
//
// NOTE:
// The caller is responsible for closing the client using client.Close().
//
// Example:
//
//	client, err := InitFirestore("my-project-id")
//	if err != nil {
//	    // handle error
//	}
func InitFirestore(projectId string) (*firestore.Client, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return client, nil
}
