package store_test

import (
	"testing"

	"git.sr.ht/~will-clarke/web-crawler/store"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	tests := []struct {
		name string
		data []string
		key  string
		want bool
	}{
		{
			name: "with a map that includes the key it can find the key",
			data: []string{"the-key"},
			key:  "the-key",
			want: true,
		},
		{
			name: "with a large map that includes the key it can find the key",
			data: []string{"aaaaaa", "another distraction", "the-key", "zzzzzz"},
			key:  "the-key",
			want: true,
		},
		{
			name: "it can't find the key if it doesn't exist",
			data: []string{"aaaaaa", "another distraction", "zzzzzz"},
			key:  "the-key",
			want: false,
		},
		{
			name: "with nothing going on it doesn't find an empty string",
			data: []string{},
			key:  "",
			want: false,
		},
		{
			name: "with nothing going on it doesn't find a specific key",
			data: []string{},
			key:  "a random key",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore()

			for _, str := range tt.data {
				s.Put(str)
			}

			ok := s.Get(tt.key)
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
				s.Put(str)
			}

			keys := s.GetAllKeys()
			assert.ElementsMatch(t, tt.want, keys)
		})
	}
}
