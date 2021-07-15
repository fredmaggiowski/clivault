package datastore

import (
	"context"
	"fmt"
)

type DataRecord []byte

type Blob map[string]DataRecord

type DataStore interface {
	LoadBlob() Blob
	WriteBlob(recordID string, record DataRecord) error
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
