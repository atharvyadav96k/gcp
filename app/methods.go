package app

import (
	"cloud.google.com/go/firestore"
)

func (a *App) Store() *firestore.Client {
	return a.FireStore.FirestoreClient
}

func (a *App) StoreDoc(collection string) *firestore.CollectionRef {
	return a.FireStore.FirestoreClient.Collection(collection)
}
