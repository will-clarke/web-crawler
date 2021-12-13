package store_test

import (
	"testing"

	"git.sr.ht/~will-clarke/web-crawler/store"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	tests := []struct {
		name string
		data map[string]bool
		key  string
		want bool
	}{
		{
			name: "with a map that includes the key it can find the key",
			data: map[string]bool{
				"the-key": true,
			},
			key:  "the-key",
			want: true,
		},
		{
			name: "with a large map that includes the key it can find the key",
			data: map[string]bool{
				"aaaaaa":              true,
				"another distraction": true,
				"the-key":             true,
				"zzzzzz":              true,
			},
			key:  "the-key",
			want: true,
		},
		{
			name: "it can't find the key if it doesn't exist",
			data: map[string]bool{
				"aaaaaa":              true,
				"another distraction": true,
				"zzzzzz":              true,
			},
			key:  "the-key",
			want: false,
		},
		{
			name: "with nothing going on it doesn't find an empty string",
			data: map[string]bool{},
			key:  "",
			want: false,
		},
		{
			name: "with nothing going on it doesn't find a specific key",
			data: map[string]bool{},
			key:  "a random key",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore()

			for k, _ := range tt.data {
				s.Put(k)
			}

			ok := s.Get(tt.key)
			assert.Equal(t, tt.want, ok)
		})
	}
}
