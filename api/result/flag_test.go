package result

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"open-feature/api/model"
	"open-feature/database"
	"testing"
)

func TestFlag_Result(t *testing.T) {
	tests := []struct {
		name    string
		input   database.Flag
		want    Flag
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"should map with one static rule",
			database.Flag{
				Key: "always one",
				Rules: []database.Rule{
					database.StaticRule(true),
				},
			},
			Flag{
				"always one",
				[]Rule{
					{
						model.RuleTypeStatic,
						true,
					},
				},
			},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Flag{}
			err := f.Result(tt.input)
			assert.Equal(t, &tt.want, f)
			tt.wantErr(t, err, fmt.Sprintf("Result(%v)", tt.input))
		})
	}
}
