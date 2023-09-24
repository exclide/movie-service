package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer_HandleRoot(t *testing.T) {
	serv := NewServer(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	serv.Root(rec, req)
	assert.Equal(t, "hello world", rec.Body.String())
}
