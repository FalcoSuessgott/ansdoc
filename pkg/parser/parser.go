package parser

import (
	"fmt"
	"strings"

	"github.com/FalcoSuessgott/ansdoc/pkg/fs"
	"gopkg.in/yaml.v3"
)

type Variable struct {
	Name        string
	Description string
	Value       interface{}
}

func ParseVars(path string) ([]*Variable, error) {
	vars := []*Variable{}

	out := fs.ReadFile(path)

	var n yaml.Node

	err := yaml.Unmarshal(out, &n)
	if err != nil {
		return nil, err
	}

	// iterate over all childs from root node
	for i, v := range n.Content[0].Content {
		// we process column bases, so we skip even entries, since these are values
		if i%2 != 0 {
			continue
		}

		// Decode varibles value
		var s interface{}

		if err := n.Content[0].Content[i+1].Decode(&s); err != nil {
			return nil, err
		}

		// remove prefix from comment
		comment := strings.TrimPrefix(v.HeadComment, "# ")

		vars = append(vars, &Variable{
			Name:        v.Value,
			Description: comment,
			Value:       fmt.Sprintf("%v", s),
		})
	}

	return vars, nil
}
