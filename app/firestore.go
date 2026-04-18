package app

import (
	"context"
)

// StoreCreate inserts a new document into the given collection with an auto-generated ID.
//
// Parameters:
//   - ctx: request context for timeout/cancellation control
//   - collection: Firestore collection name
//   - data: struct or map to be stored
//
// Returns:
//   - error if the write operation fails
//
// Example:
//
//	err := app.StoreCreate(ctx, "users", userData)
func (a *App) StoreCreate(ctx context.Context, collection string, data interface{}) error {
	_, _, err := a.StoreDoc(collection).Add(ctx, data)
	return err
}

// StoreCreateWithId creates a new document with a specific ID in the given collection.
// It fails if a document with the same ID already exists.
//
// Parameters:
//   - ctx: request context
//   - collection: Firestore collection name
//   - id: custom document ID
//   - data: struct or map to be stored
//
// Returns:
//   - error if document already exists or write fails
//
// Example:
//
//	err := app.StoreCreateWithId(ctx, "users", "user123", userData)
func (a *App) StoreCreateWithId(ctx context.Context, collection string, id string, data interface{}) error {
	docRef := a.StoreDoc(collection).Doc(id)
	_, err := docRef.Create(ctx, data)
	return err
}

// StoreUpdate updates or overwrites a document with the given ID.
// If the document does not exist, it will be created.
//
// Parameters:
//   - ctx: request context
//   - collection: Firestore collection name
//   - id: document ID
//   - data: updated data (replaces existing document)
//
// Returns:
//   - error if update fails
//
// Example:
//
//	err := app.StoreUpdate(ctx, "users", "user123", updatedData)
func (a *App) StoreUpdate(ctx context.Context, collection string, id string, data interface{}) error {
	docRef := a.StoreDoc(collection).Doc(id)
	_, err := docRef.Set(ctx, data)
	return err
}

// StoreDelete deletes a document from the given collection by ID.
//
// Parameters:
//   - ctx: request context
//   - collection: Firestore collection name
//   - id: document ID
//
// Returns:
//   - error if deletion fails
//
// Example:
//
//	err := app.StoreDelete(ctx, "users", "user123")
func (a *App) StoreDelete(ctx context.Context, collection string, id string) error {
	_, err := a.StoreDoc(collection).Doc(id).Delete(ctx)
	return err
}
