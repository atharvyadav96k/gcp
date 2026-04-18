package app

import (
	"context"
)

func (a *App) StoreCreate(ctx context.Context, collection string, data interface{}) error {
	_, _, err := a.StoreDoc(collection).Add(ctx, data)
	return err
}

func (a *App) StoreCreateWithId(ctx context.Context, collection string, id string, data interface{}) error {
	docRef := a.StoreDoc(collection).Doc(id)
	_, err := docRef.Create(ctx, data)
	return err
}

func (a *App) StoreUpdate(ctx context.Context, collection string, id string, data interface{}) error {
	docRef := a.StoreDoc(collection).Doc(id)
	_, err := docRef.Set(ctx, data)
	return err
}

func (a *App) StoreDelete(ctx context.Context, collection string, id string) error {
	_, err := a.StoreDoc(collection).Doc(id).Delete(ctx)
	return err
}
