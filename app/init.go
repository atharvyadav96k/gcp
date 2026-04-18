package app

import (
	"context"
	"fmt"

	"github.com/atharvyadav96k/gcp/app/models/firestore"
	"github.com/atharvyadav96k/gcp/app/models/pubsub"
	"github.com/atharvyadav96k/gcp/app/models/secrets"
)

// InitEnvironmentVariables loads and initializes environment variables
// and secrets required by the application.
//
// It populates the Env field using the secrets package.
//
// Example:
//
//	app.InitEnvironmentVariables()
func (a *App) InitEnvironmentVariables() {
	a.Env = *secrets.NewSecrets()
}

// Init initializes the App dependencies (Firestore and PubSub).
//
// NOTE:
// This function RETURNS a new App instance. Make sure you use the returned value,
// otherwise initialization will be lost.
//
// Example:
//
//	app := new(App).Init()
func (a *App) Init() *App {
	return &App{
		FireStore: &firestore.Service{},
		PubSub:    &pubsub.Service{},
	}
}

// InitFirestore initializes the Firestore client using the given project ID.
// It ensures initialization happens only once using sync.Once.
//
// Parameters:
//   - projectId: Google Cloud project ID
//
// Returns:
//   - error if initialization fails
//
// Example:
//
//	err := app.InitFirestore("my-project-id")
func (a *App) InitFirestore(projectId string) error {
	if a.FireStore == nil {
		a.FireStore = &firestore.Service{}
	}

	var err error

	a.FireStore.Once.Do(func() {
		client, initErr := firestore.InitFirestore(projectId)
		if initErr != nil {
			err = initErr
			return
		}
		a.FireStore.Client = client
	})

	return err
}

// InitPubSub initializes the Pub/Sub client using the given project ID.
// It ensures initialization happens only once using sync.Once.
//
// Parameters:
//   - projectId: Google Cloud project ID
//
// Returns:
//   - error if initialization fails
//
// Example:
//
//	err := app.InitPubSub("my-project-id")
func (a *App) InitPubSub(ctx context.Context, projectId string) error {
	if a.PubSub == nil {
		a.PubSub = &pubsub.Service{}
	}

	// already initialized
	if a.PubSub.Client != nil {
		return nil
	}

	client, err := pubsub.InitPubSub(ctx, projectId)
	if err != nil {
		return err
	}

	a.PubSub.Client = client
	return nil
}

// Close gracefully closes all initialized external service clients
// such as Firestore and Pub/Sub.
//
// It should be called during application shutdown to release resources.
//
// Example:
//
//	defer app.Close()
func (a *App) Close() {
	if a.FireStore != nil {
		if err := a.FireStore.Close(); err != nil {
			fmt.Printf("Error closing Firestore client: %v\n", err)
		}
	}

	if a.PubSub != nil {
		if err := a.PubSub.Close(); err != nil {
			fmt.Printf("Error closing PubSub client: %v\n", err)
		}
	}
}
