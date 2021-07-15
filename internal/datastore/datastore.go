package datastore

import (
	"context"
	"fmt"
)

type DataRecord string

type DataStore interface {
	LoadBlob(blob []byte)
	WriteBlob(recordID string, record DataRecord) ([]byte, error)
}

type dataStoreContextKey struct{}

func WithValue(ctx context.Context, ds DataStore) context.Context {
	return context.WithValue(ctx, dataStoreContextKey{}, ds)
}

func FromContext(ctx context.Context) (DataStore, error) {
	ds, ok := ctx.Value(dataStoreContextKey{}).(DataStore)
	if !ok {
		return nil, fmt.Errorf("no DataStore in context")
	}
	return ds, nil
}
