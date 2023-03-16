package writer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSplitFile(t *testing.T) {
	testCases := []struct {
		name      string
		start     string
		end       string
		delimiter string
		err       bool
	}{
		{
			name: "simple",
			start: `1
2
3
`,
			end: `
6
7
8
9`,
			delimiter: Delimiter,
		},
		{
			name:  "simple_2",
			start: ``,
			end: `
7
8
9`,
			delimiter: Delimiter,
		},
		{
			name:      "invalid",
			err:       true,
			delimiter: Delimiter,
		},
		{
			name:      "invalid_2",
			err:       true,
			delimiter: Delimiter,
		},
		{
			name:      "invalid_3",
			err:       true,
			delimiter: Delimiter,
		},
	}

	for _, tc := range testCases {
		start, end, err := SplitFile(fmt.Sprintf("testdata/%s.txt", tc.name), tc.delimiter)

		if tc.err {
			require.Error(t, err, tc.name)

			continue
		}

		assert.NoError(t, err, tc.name)
		assert.Equal(t, tc.start, start, tc.name)
		assert.Equal(t, tc.end, end, tc.name)
	}
}
