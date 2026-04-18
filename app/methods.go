package app

import "cloud.google.com/go/firestore"

func (a *App) Database() *firestore.Client {
	return a.FireStore.FirestoreClient
}
