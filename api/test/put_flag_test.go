package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"open-feature/database"
	"strings"
	"testing"
)

func Test_PutFlag_Should400_WhenInputIsMissingFields(t *testing.T) {
	s := initializeTestServer()

	resp, err := s.Post(
		"/flags/",
		"application/json",
		strings.NewReader(`{"key":"","rules":[{"type":"static","data":true}]}`),
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_PutFlag_ShouldInsertNewFlag_WhenInputIsAllValid(t *testing.T) {
	s := initializeTestServer()

	resp, err := s.Post(
		"/flags/",
		"application/json",
		strings.NewReader(`{"key":"test-flag","rules":[{"type":"static","data":true}]}`),
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	flag, _ := s.api.Get("test-flag")
	assert.Equal(t, "test-flag", flag.Key)
	assert.True(t, flag.Rules[0].Eval())
}

func Test_PutFlag_ShouldUpdateExistingFlag_WhenInputIsAllValid(t *testing.T) {
	s := initializeTestServer()
	_ = s.api.Put(database.Flag{
		Key: "test-flag",
		Rules: []database.Rule{
			database.StaticRule(false),
		},
	})

	resp, err := s.Post(
		"/flags/",
		"application/json",
		strings.NewReader(`{"key":"test-flag","rules":[{"type":"static","data":true}]}`),
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	flag, _ := s.api.Get("test-flag")
	assert.True(t, flag.Rules[0].Eval())
}
