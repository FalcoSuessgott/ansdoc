package parser

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const tmpl = `<table>
<tr>
<th>Name</th>
<th>Description</th>
<th>Default Value</th>
</tr>
{{ range .}}
<tr>
<td>` +
	"\n\n" +
	`{{ .Name }}` +
	"\n\n" +
	`</td>` +
	`<td>{{ .Description }}</td>` +
	`<td>` +
	"\n\n```yaml\n" +
	`{{ .Value }}` +
	"\n```\n\n" +
	`</td>` +
	`</tr>
{{ end }}
</table>
`

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

	defer f.Close()

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

		d, err := yaml.Marshal(s)
		if err != nil {
			return nil, fmt.Errorf("error while marshalling %s to yaml: %w", v.Value, err)
		}

		vars = append(vars, &Variable{
			Name:        v.Value,
			Description: trimPrefix(v.HeadComment),
			Value:       string(d), // todo: fix this
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

// Render takes a slice of variables and renders those in a html table with its values in a code block.
func Render(vars []*Variable) ([]byte, error) {
	// in order to support multiline code blocks we render a html table
	var buf bytes.Buffer

	tpl, err := template.New("template").Option("missingkey=error").Parse(string(tmpl))
	if err != nil {
		return buf.Bytes(), err
	}

	if err := tpl.Execute(&buf, vars); err != nil {
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
}
