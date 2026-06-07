package storage

import (
	"context"
	"sync"
)

// MemoryEngine is an in-memory implementation of the StorageEngine interface.
type MemoryEngine struct {
	mu   sync.RWMutex      // Mutex to ensure thread-safe access to the data map.
	data map[string][]byte // Map to store key-value pairs in memory.
}

// NewMemoryEngine initializes and returns a new MemoryEngine instance.
func NewMemoryEngine() *MemoryEngine {
	return &MemoryEngine{
		data: make(map[string][]byte), // Initialize the data map to store key-value pairs.
	}
}

// Put stores a key-value pair in the memory engine.
func (m *MemoryEngine) Put(ctx context.Context, key string, value []byte) error {
	// Validate the key before storing.
	if key == "" {
		return ErrEmptyKey
	}

	// Lock the mutex to ensure thread-safe access to the data map.
	m.mu.Lock()
	// Unlock the mutex after the function returns to allow other operations to proceed.
	defer m.mu.Unlock()

	// Store the key-value pair in the data map.
	m.data[key] = value
	// Return nil to indicate that the operation was successful.
	return nil
}

// Get retrieves the value associated with the given key from the memory engine.
func (m *MemoryEngine) Get(ctx context.Context, key string) ([]byte, error) {
	// Validate the key before retrieving.
	if key == "" {
		return nil, ErrEmptyKey
	}

	// Lock the mutex for reading to ensure thread-safe access to the data map.
	m.mu.RLock()
	// Unlock the mutex after the function returns to allow other operations to proceed.
	defer m.mu.RUnlock()

	// Retrieve the value associated with the key from the data map.
	value, exists := m.data[key]
	// If the key does not exist in the data map, return an error indicating
	// that the key was not found.
	if !exists {
		return nil, ErrKeyNotFound
	}
	// Return the retrieved value and nil to indicate that the operation was successful.
	return value, nil
}

// Delete removes a key-value pair from the memory engine based on the given key.
func (m *MemoryEngine) Delete(ctx context.Context, key string) error {
	// Validate the key before deleting.
	if key == "" {
		return ErrEmptyKey
	}

	// Lock the mutex to ensure thread-safe access to the data map.
	m.mu.Lock()
	// Unlock the mutex after the function returns to allow other operations to proceed.
	defer m.mu.Unlock()

	// Check if the key exists in the data map before attempting to delete it.
	if _, exists := m.data[key]; !exists {
		return ErrKeyNotFound
	}
	// Delete the key-value pair from the data map.
	delete(m.data, key)
	// Return nil to indicate that the operation was successful.
	return nil
}

// Close performs any necessary cleanup for the memory engine.
// Since this is an in-memory implementation,
// there are no resources to clean up, so it simply returns nil.
func (m *MemoryEngine) Close() error {
	return nil
}
