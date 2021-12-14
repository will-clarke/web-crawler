package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"git.sr.ht/~will-clarke/web-crawler/server"
	"git.sr.ht/~will-clarke/web-crawler/store"
)

func TestPingRoute(t *testing.T) {
	router := server.SetupRouter(store.NewStore())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}
