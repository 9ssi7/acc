// Package acc provides an accumulator mechanism to accumulate, process, and manage data.
package acc

import "context"

// DataStorage defines the interface for any storage mechanism to save and load data.
type DataStorage[T interface{}] interface {
	// Save stores the provided data items.
	Save(ctx context.Context, data []T) error
	// Load retrieves the stored data items.
	Load(ctx context.Context) ([]T, error)
}

// InMemoryStorage is an in-memory storage implementation for the DataStorage interface.
type InMemoryStorage[T any] struct {
	data []Data[T]
}

// NewInMemoryStorage initializes a new instance of InMemoryStorage.
func NewInMemoryStorage[T any]() *InMemoryStorage[T] {
	return &InMemoryStorage[T]{}
}

// Save stores the provided data items in memory.
func (s *InMemoryStorage[T]) Save(ctx context.Context, data []Data[T]) error {
	s.data = append(s.data, data...)
	return nil
}

// Load retrieves the stored data items from memory.
func (s *InMemoryStorage[T]) Load(ctx context.Context) ([]Data[T], error) {
	return s.data, nil
}
