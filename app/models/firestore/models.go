package firestore

import (
	"sync"

	"cloud.google.com/go/firestore"
)

// Service wraps the Firestore client and manages its lifecycle.
//
// It ensures the client is initialized only once using sync.Once.
// This struct should be reused across the application (singleton-style)
// to avoid multiple client initializations.
type Service struct {

	// Client is the underlying Firestore client used to interact
	// with the database (collections, documents, queries, etc.).
	Client *firestore.Client

	// once guarantees that the Firestore client is initialized only once,
	// even if Init is called multiple times from different goroutines.
	Once sync.Once
}
