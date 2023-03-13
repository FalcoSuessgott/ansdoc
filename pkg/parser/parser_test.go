package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseConfig(t *testing.T) {
	testCases := []struct {
		name string
		exp  []*Variable
		err  bool
	}{
		{
			name: "simple vars",
			exp: []*Variable{
				{
					Name:        "var1",
					Value:       true,
					Description: "# bool",
				},
				{
					Name:        "var2",
					Value:       "string",
					Description: "# string",
				},
				{
					Name:        "var3",
					Value:       42,
					Description: "# int",
				},
				{
					Name:        "var4",
					Value:       []interface{}{1, 2, 3},
					Description: "# list",
				},
				{
					Name: "var5",
					Value: map[string]interface{}{
						"linux":   true,
						"mac":     false,
						"windows": false,
					},
					Description: "# map",
				},
				{
					Name:        "var6",
					Value:       "value with a multiline comment",
					Description: "# multiline \n# comment",
				},
			},
		},
	}

	for i, tc := range testCases {
		cfg, err := ParseVars(fmt.Sprintf("testdata/%v.yml", i+1))
		if tc.err {
			require.Error(t, err, tc.name)
		}

		assert.EqualValues(t, tc.exp, cfg, tc.name)
	}
}
