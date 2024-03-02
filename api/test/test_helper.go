package test

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"open-feature/api"
	"open-feature/database"
	"testing"
)

type testServer struct {
	httptest *httptest.Server
	api      *api.Server
}

func (t testServer) Get(url string) (*http.Response, error) {
	return t.httptest.Client().Get(t.httptest.URL + url)
}

func (t testServer) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	return t.httptest.Client().Post(t.httptest.URL+url, contentType, body)
}

func initializeTestServer() testServer {
	s := api.Server{
		Database: database.NewMemDatabase(),
	}

	return testServer{
		httptest.NewServer(s.NewServeMux()),
		&s,
	}
}

func assertJSONBody(t *testing.T, resp *http.Response, jsonExpected string) {
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, jsonExpected, string(body))
}
