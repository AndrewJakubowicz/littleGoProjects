package mockingTest1

import "sync"

type inMemoryHashTable struct {
	m   map[string][]byte
	lck sync.RWMutex
}

// NewInMemoryHashTable creates a new in memory hash table.
// This is used in mocking tests as it shares the same interface
// as simple key value stores.
func NewInMemoryHashTable() HashTable {
	return &inMemoryHashTable{m: make(map[string][]byte)}
}

func (i *inMemoryHashTable) Get(key string) ([]byte, error) {
	i.lck.RLock() // lock for reading
	defer i.lck.RUnlock()
	val, ok := i.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

func (i *inMemoryHashTable) Set(key string, value []byte) error {
	i.lck.Lock() // lock for writing
	defer i.lck.Unlock()
	i.m[key] = value
	return nil
}
