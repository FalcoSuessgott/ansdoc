package cmd

import (
	"fmt"
	"os"

	"github.com/FalcoSuessgott/ansdoc/pkg/parser"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var file = "defaults/main.yml"

func newRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "ansdoc",
		Short:         "out of the box documentation for you ansible roles",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			vars, err := parser.ParseVars(file)
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"variable", "description", "default value"})
			table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
			table.SetCenterSeparator("|")

			data := [][]string{}

			for _, v := range vars {
				data = append(data, []string{fmt.Sprintf("`%s`", v.Name), v.Description, fmt.Sprintf("`%v`", v.Value)})
			}

			table.SetAutoWrapText(false)
			table.AppendBulk(data)
			table.Render()

			return nil
		},
	}

	cmd.AddCommand(versionCmd(version))

	cmd.Flags().StringVarP(&file, "file", "f", file, "path to the variables file")

	return cmd
}

// Execute invokes the command.
func Execute(version string) error {
	if err := newRootCmd(version).Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}
