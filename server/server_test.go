package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"git.sr.ht/~will-clarke/web-crawler/server"
	"git.sr.ht/~will-clarke/web-crawler/store"
)

func TestHealthcheck(t *testing.T) {
	router := server.SetupRouter(store.NewStore())

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}

func TestCrawlPost(t *testing.T) {
	router := server.SetupRouter(store.NewStore())

	w := httptest.NewRecorder()
	body := `{"url":"fake-url"}`
	req, err := http.NewRequest("POST", "/crawl", strings.NewReader(body))
	req.SetBasicAuth(server.TopSecretUsername, server.TopSecretPassword)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCrawlGet(t *testing.T) {
	router := server.SetupRouter(store.NewStore())

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/crawl/some-id", nil)
	req.SetBasicAuth(server.TopSecretUsername, server.TopSecretPassword)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":"some-id","links":[]}`, w.Body.String())
}
