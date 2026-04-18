package app

import (
	"cloud.google.com/go/firestore"
)

// Store returns the initialized Firestore client.
//
// Make sure InitFirestore() has been called before using this,
// otherwise it may return nil.
//
// Example:
//
//	client := app.Store()
func (a *App) Store() *firestore.Client {
	if a.FireStore == nil {
		return nil
	}
	return a.FireStore.Client
}

// StoreDoc returns a reference to a Firestore collection.
//
// Parameters:
//   - collection: Firestore collection name
//
// Returns:
//   - *firestore.CollectionRef for performing operations like Add, Doc, Where, etc.
//   - nil if Firestore is not initialized
//
// Example:
//
//	users := app.StoreDoc("users")
//	users.Doc("id").Set(ctx, data)
func (a *App) StoreDoc(collection string) *firestore.CollectionRef {
	client := a.Store()
	if client == nil {
		return nil
	}
	return client.Collection(collection)
}
