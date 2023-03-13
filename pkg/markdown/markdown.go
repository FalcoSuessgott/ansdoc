package markdown

import (
	"fmt"
	"os"

	"github.com/FalcoSuessgott/ansdoc/pkg/parser"
	"github.com/olekukonko/tablewriter"
)

// NewMarkdownTable builds a new markdown table.
func NewMarkdownTable(vars []*parser.Variable) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"variable", "description", "default value"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	data := [][]string{}

	for _, v := range vars {
		data = append(data, []string{fmt.Sprintf("`%s`", v.Name), v.Description, fmt.Sprintf("`%#v`", v.Value)})
	}

	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()

	return nil
}
