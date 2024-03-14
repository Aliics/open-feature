package input

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestGetValidatedInput_Flag(t *testing.T) {
	type testCase[V Validator] struct {
		name    string
		rawJSON string
		want    *V
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[Flag]{
		{
			"should fail if missing key",
			`{"key":"","rules":[{"type":"static","data":true}]}`,
			nil,
			assert.Error,
		},
		{
			"should fail if missing rules",
			`{"key":"awesome-key","rules":[]}`,
			nil,
			assert.Error,
		},
		{
			"should fail if missing rules",
			`{"key":"awesome-key","rules":[{"type":"static","data":true}]}`,
			&Flag{
				"awesome-key",
				[]Rule{
					{"static", true},
				},
			},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("nop", http.MethodPost, strings.NewReader(tt.rawJSON))
			assert.NoError(t, err)

			got, err := GetValidatedInput[Flag](req)
			if !tt.wantErr(t, err, fmt.Sprintf("GetValidatedInput(%v)", tt.rawJSON)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetValidatedInput(%v)", tt.rawJSON)
		})
	}
}
