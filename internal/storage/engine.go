package storage

import (
	"context"
	"errors"
)

var (
	ErrKeyNotFound = errors.New("key not found in storage engine")
	ErrEmptyKey    = errors.New("key cannot be empty")
)

// StorageEngine defines the transactional capabilities of our local and distributed storage.
type StorageEngine interface {
	Put(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
	Close() error
}
