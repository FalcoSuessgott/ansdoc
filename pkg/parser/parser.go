package parser

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Variable holds a variable including its description.
type Variable struct {
	Name        string
	Description string
	Value       interface{}
}

// ParseVars parses a vars file.
func ParseVars(path string) ([]*Variable, error) {
	vars := []*Variable{}

	var n yaml.Node

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	out, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(out, &n); err != nil {
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

		vars = append(vars, &Variable{
			Name:        v.Value,
			Description: trimPrefix(v.HeadComment),
			Value:       s,
		})
	}

	return vars, nil
}

// don't ask ...
func trimPrefix(s string) string {
	res := ""

	s = strings.ReplaceAll(s, "#", "")

	for _, p := range strings.Split(s, "\n") {
		res += fmt.Sprintf("%s ", strings.TrimSpace(p))
	}

	return strings.TrimSpace(res)
}
