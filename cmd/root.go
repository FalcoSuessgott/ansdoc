package cmd

import (
	"fmt"

	"github.com/FalcoSuessgott/ansdoc/pkg/markdown"
	"github.com/FalcoSuessgott/ansdoc/pkg/parser"
	"github.com/spf13/cobra"
)

var file = "defaults/main.yml"

func newRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ansdoc",
		Short: "out of the box documentation for you ansible roles",
		RunE: func(cmd *cobra.Command, args []string) error {
			vars, err := parser.ParseVars(file)
			if err != nil {
				return err
			}

			if markdown.NewMarkdownTable(vars) != nil {
				return err
			}

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
