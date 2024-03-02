package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"open-feature/database"
	"testing"
)

func Test_GetFlag_Should404_WhenThereIsNoFlagWithKey(t *testing.T) {
	s := initializeTestServer()

	resp, err := s.Get("/flags/test-flag")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func Test_GetFlag_ShouldRespondWithFlag_WhenFlagExistsWithKey(t *testing.T) {
	s := initializeTestServer()
	_ = s.api.Put(database.Flag{
		Key: "test-flag",
		Rules: []database.Rule{
			database.StaticRule(true),
		},
	})

	resp, err := s.Get("/flags/test-flag")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assertJSONBody(t, resp, `{"key":"test-flag","rules":[{"type":"static","data":true}]}`)
}
