package result

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"open-feature/database"
	"testing"
)

func TestFlag_Map(t *testing.T) {
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
						RuleTypeStatic,
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
			err := f.Map(tt.input)
			assert.Equal(t, &tt.want, f)
			tt.wantErr(t, err, fmt.Sprintf("Map(%v)", tt.input))
		})
	}
}
