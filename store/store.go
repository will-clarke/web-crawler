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

// Exists will return whether an item exists stored first in
// the (top-level) map, then in the second map.
func (s *Store) Exists(namespaceKey, key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	mapForNamespace, exists := s.data[namespaceKey]
	if !exists {
		return false
	}
	_, exists = mapForNamespace[key]
	return exists
}

// Put will update the store;  the specific "namespaceKey" / "key" combination will now `Exist`.
func (s *Store) Put(namespaceKey, key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mapForNamespace, exists := s.data[namespaceKey]
	if !exists {
		s.data[namespaceKey] = map[string]bool{
			key: true,
		}
		return
	}
	mapForNamespace[key] = true
}

// GetAllKeys returns all keys for a given namespace.
func (s *Store) GetAllKeys(namespaceKey string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	keys := []string{}
	for k := range s.data[namespaceKey] {
		keys = append(keys, k)
	}
	return keys
}
