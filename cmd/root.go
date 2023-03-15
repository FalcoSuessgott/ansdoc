package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/FalcoSuessgott/ansdoc/pkg/parser"
	"github.com/caarlos0/env/v6"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Opts contains all cli args.
type Opts struct {
	File       string `env:"FILE" envDefault:"defaults/main.yml"`
	OutputFile string `env:"OUTPUT_FILE"`
	Backup     bool   `env:"BACKUP"`
	Insert     bool   `env:"INSERT"`
}

func newRootCmd(version string) *cobra.Command {
	opts := &Opts{}

	if err := env.Parse(opts, env.Options{
		Prefix: "ANSDOC_",
	}); err != nil {
		log.Fatal(err)
	}

	cmd := &cobra.Command{
		Use:           "ansdoc",
		Short:         "out of the box documentation for you ansible roles",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(opts)

			vars, err := parser.ParseVars(opts.File)
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

	cmd.Flags().StringVarP(&opts.File, "file", "f", opts.File, "path to the variables file")
	cmd.Flags().StringVarP(&opts.OutputFile, "output-file", "o", opts.OutputFile, "where to write the output to (required insert mode)")
	cmd.Flags().BoolVarP(&opts.Backup, "backup", "b", opts.Backup, "backup the output file before writing")
	cmd.Flags().BoolVarP(&opts.Insert, "insert", "i", opts.Insert, "insert mode, inserts the markdown table in the specified output file")

	return cmd
}

// Execute invokes the command.
func Execute(version string) error {
	if err := newRootCmd(version).Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}
