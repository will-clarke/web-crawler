package store_test

import (
	"testing"

	"git.sr.ht/~will-clarke/web-crawler/store"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	tests := []struct {
		name      string
		data      map[string]map[string]bool
		namespace string
		key       string
		want      bool
	}{
		{
			name: "with a map that includes the key it can find the key",
			data: map[string]map[string]bool{
				"aaa": {
					"the-key": true,
				},
			},
			namespace: "aaa",
			key:       "the-key",
			want:      true,
		},
		{
			name: "with a large map that includes the key it can find the key",
			data: map[string]map[string]bool{
				"aaa": {
					"aaaaaaaa":            true,
					"the-key":             true,
					"another distraction": true,
				},

				"zzz": {
					"zzzzzzzzz": true,
				},
			},
			namespace: "aaa",
			key:       "the-key",
			want:      true,
		},
		{
			name: "it can't find the key if it doesn't exist",
			data: map[string]map[string]bool{
				"aaa": {
					"aaaaaaaa":            true,
					"another distraction": true,
				},

				"zzz": {
					"zzzzzzzzz": true,
				},
			},
			namespace: "aaa",
			key:       "the-key",
			want:      false,
		},
		{
			name: "with a large map that includes the key in a different namespace",
			data: map[string]map[string]bool{
				"aaa": {
					"aaaaaaaa":            true,
					"the-key":             true,
					"another distraction": true,
				},

				"zzz": {
					"zzzzzzzzz": true,
				},
			},
			namespace: "zzz",
			key:       "the-key",
			want:      false,
		},
		{
			name:      "with no namespace or key it can't find an empty string",
			data:      map[string]map[string]bool{},
			namespace: "",
			key:       "",
			want:      false,
		},
		{
			name:      "with no key on it doesn't find an empty string",
			data:      map[string]map[string]bool{},
			namespace: "aaa",
			key:       "",
			want:      false,
		},
		{
			name:      "with no key on it doesn't find an empty string",
			data:      map[string]map[string]bool{},
			namespace: "aaa",
			key:       "this doesn't exist either",
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore()

			for namespace, namespaceMap := range tt.data {
				for key := range namespaceMap {
					s.Put(namespace, key)
				}
			}

			ok := s.Exists(tt.namespace, tt.key)
			assert.Equal(t, tt.want, ok)
		})
	}
}

func TestStore_GetAllKeys(t *testing.T) {
	tests := []struct {
		name string
		data []string
		want []string
	}{
		{
			name: "with no keys",
			data: []string{},
			want: []string{},
		},
		{
			name: "with one key",
			data: []string{"a"},
			want: []string{"a"},
		},
		{
			name: "with four keys",
			data: []string{"aaaaaa", "bbbbb", "cccccc", "zzzzzz"},
			want: []string{"aaaaaa", "bbbbb", "cccccc", "zzzzzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore()

			for _, str := range tt.data {
				s.Put("a", str)
			}

			keys := s.GetAllKeys("a")
			assert.ElementsMatch(t, tt.want, keys)
			keysForWrongNamespace := s.GetAllKeys("z")
			assert.Empty(t, keysForWrongNamespace)
		})
	}
}
