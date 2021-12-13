package store

import (
	"sync"
)

type Store struct {
	mu   sync.Mutex
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, exists := s.data[key]
	return val, exists
}

func (s *Store) Put(k, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[k] = v
}
