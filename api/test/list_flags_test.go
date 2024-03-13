package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"open-feature/database"
	"testing"
)

func Test_FlagsAll_ShouldReturnEmptyList_WhenNoFlagsHaveBeenCreated(t *testing.T) {
	s := initializeTestServer()

	resp, err := s.Get("/flags/")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assertJSONBody(t, resp, `[]`)
}

func Test_FlagsAll_ShouldReturnOneFlag_WhenAFlagHasBeenInserted(t *testing.T) {
	s := initializeTestServer()
	_ = s.api.Database.Put(database.Flag{
		Key: "nice-test-flag",
		Rules: []database.Rule{
			database.StaticRule(true),
		},
	})

	resp, err := s.Get("/flags/")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assertJSONBody(
		t,
		resp,
		`[
		  {
			"key": "nice-test-flag",
			"rules": [
			  {
				"type": "static",
				"data": true
			  }
			]
		  }
		]`,
	)
}
