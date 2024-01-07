package inmemorystore

import (
	"sync"
)

type inMemoryStore struct {
	mu   sync.RWMutex
	data [][]byte
}

func NewInMemoryStore() *inMemoryStore {
	return &inMemoryStore{
		data: make([][]byte, 0),
	}
}

func (i *inMemoryStore) Push(input []byte) (int, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.data = append(i.data, input)

	return len(i.data) - 1, nil
}

func (i *inMemoryStore) Get(input int) ([]byte, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return i.data[input], nil
}
