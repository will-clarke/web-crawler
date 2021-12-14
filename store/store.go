package store

import (
	"sync"
)

type namespace = string
type key = string

type Store struct {
	mu   sync.Mutex
	data map[namespace]map[key]bool
	// this type isn't that scary!
	// We've basically just got a list of 'namespaces' (an example could
	// be a collection of webpages we've scraped from a single job)
	// which themselves have a list of keys (eg. the individual URLs visited).
	//
	// Also, I could technically be slightly more efficient
	// (in theory) by using an 'interface{}' here instead of a bool...
	// but I don't think it's worth the extra complexity here.
}

func NewStore() *Store {
	return &Store{
		data: make(map[namespace]map[key]bool),
	}
}

func (s *Store) Get(namespace, key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	mapForNamespace, exists := s.data[namespace]
	if !exists {
		return false
	}
	_, exists = mapForNamespace[key]
	return exists
}

func (s *Store) Put(namespace, key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mapForNamespace, exists := s.data[namespace]
	if !exists {
		s.data[namespace] = map[string]bool{
			key: true,
		}
		return
	}
	mapForNamespace[key] = true
}

func (s *Store) GetAllKeys(namespace string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	keys := []string{}
	for k := range s.data[namespace] {
		keys = append(keys, k)
	}
	return keys
}
