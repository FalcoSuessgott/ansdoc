package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrimPrefix(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
	}{
		{
			input:  "# simple",
			expect: "simple",
		},
		{
			input:  "# multiline \n# comment",
			expect: "multiline comment",
		},
		{
			input:  "#nospace",
			expect: "nospace",
		},
		{
			input:  "#multiline\n#nospace",
			expect: "multiline nospace",
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expect, trimPrefix(tc.input))
	}
}

//nolint: funlen
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
					Description: "bool",
				},
				{
					Name:        "var2",
					Value:       "string",
					Description: "string",
				},
				{
					Name:        "var3",
					Value:       42,
					Description: "int",
				},
				{
					Name:        "var4",
					Value:       []interface{}{1, 2, 3},
					Description: "list",
				},
				{
					Name: "var5",
					Value: map[string]interface{}{
						"linux":   true,
						"mac":     false,
						"windows": false,
					},
					Description: "map",
				},
				{
					Name:        "var6",
					Value:       "value with a multiline comment",
					Description: "multiline comment",
				},
				{
					Name:  "var7",
					Value: "no description",
				},
				{
					Name: "var8",
					Value: `[
{% for server in groups[vault_raft_group_name] %}
  {
    "peer": "{{ server }}",
    "api_addr": "{{ hostvars[server]['vault_api_addr'] |
    default(vault_protocol + '://' + hostvars[server]['ansible_' + hostvars[server]['ansible_default_ipv4']['interface']]['ipv4']['address'] + ':' + (vault_port|string)) }}"
  },
{% endfor %}
]`,
					Description: "jinja",
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
