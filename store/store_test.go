package store_test

import (
	"testing"

	"git.sr.ht/~will-clarke/web-crawler/store"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]string
		key     string
		wantStr string
		wantOk  bool
	}{
		{
			name: "with a map that includes the key it can find the key",
			data: map[string]string{
				"the-key": "the-value",
			},
			key:     "the-key",
			wantStr: "the-value",
			wantOk:  true,
		},
		{
			name: "with a large map that includes the key it can find the key",
			data: map[string]string{
				"aaaaaa":              "nope",
				"another distraction": "nope",
				"the-key":             "the-value",
				"zzzzzz":              "nope",
			},
			key:     "the-key",
			wantStr: "the-value",
			wantOk:  true,
		},
		{
			name: "it can't find the key if it doesn't exist",
			data: map[string]string{
				"aaaaaa":              "nope",
				"another distraction": "nope",
				"zzzzzz":              "nope",
			},
			key:     "the-key",
			wantStr: "",
			wantOk:  false,
		},
		{
			name:    "with nothing going on it doesn't find an empty string",
			data:    map[string]string{},
			key:     "",
			wantStr: "",
			wantOk:  false,
		},
		{
			name:    "with nothing going on it doesn't find a specific key",
			data:    map[string]string{},
			key:     "a random key",
			wantStr: "",
			wantOk:  false,
		},
		{
			name: "with a map that includes the key it can find the key",
			data: map[string]string{
				"the-key": "the-value",
			},
			key:     "the-key",
			wantStr: "the-value",
			wantOk:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore()

			for k, v := range tt.data {
				s.Put(k, v)
			}

			str, ok := s.Get(tt.key)
			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.wantStr, str)
		})
	}
}
