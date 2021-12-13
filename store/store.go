package store

import (
	"sync"
)

type Store struct {
	mu   sync.Mutex
	data map[string]bool
	// Technically I could be 'clever' and be slightly more efficient
	// (in theory) by using an 'interface{}' here instead of a bool...
	// but I don't think it's worth the extra complexity here.
}

func NewStore() Store {
	return Store{
		data: make(map[string]bool),
	}
}

func (s *Store) Get(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[key]
	return exists
}

func (s *Store) Put(k string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[k] = true
}

func (s *Store) GetAllKeys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	keys := []string{}
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}
