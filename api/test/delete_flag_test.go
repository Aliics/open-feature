package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"open-feature/database"
	"testing"
)

func Test_DeleteFlag_ShouldRespondWithNotFound_WhenNoFlagExists(t *testing.T) {
	s := initializeTestServer()

	resp, err := s.Delete("/flags/what-flag")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func Test_DeleteFlag_ShouldRespondOK_WhenAFlagExists(t *testing.T) {
	s := initializeTestServer()
	_ = s.api.Database.Put(database.Flag{
		Key: "nice-test-flag",
		Rules: []database.Rule{
			database.StaticRule(true),
		},
	})

	resp, err := s.Get("/flags/nice-test-flag")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
