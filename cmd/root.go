package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FalcoSuessgott/ansdoc/pkg/parser"
	"github.com/FalcoSuessgott/ansdoc/pkg/writer"
	"github.com/caarlos0/env/v6"
	"github.com/spf13/cobra"
)

// Opts contains all cli args.
type Opts struct {
	File       string `env:"FILE" envDefault:"defaults/main.yml"`
	Markdown   bool   `env:"MARKDOWN" envDefault:"false"`
	OutputFile string `env:"OUTPUT_FILE"`
	Backup     bool   `env:"BACKUP"`
	Insert     bool   `env:"INSERT"`
}

// nolint: funlen, cyclop
func newRootCmd(version string) *cobra.Command {
	opts := &Opts{}

	if err := env.Parse(opts, env.Options{
		Prefix: "ANSDOC_",
	}); err != nil {
		log.Fatal(err)
	}

	cmd := &cobra.Command{
		Use:           "ansdoc",
		Short:         "out-of-the-box documentation for you ansible roles",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Insert && opts.OutputFile == "" {
				return fmt.Errorf("when using insert mode you need to specify an output file using --outputf-file / -o")
			}

			vars, err := parser.ParseVars(opts.File)
			if err != nil {
				return err
			}

			fmt.Printf("read variables from %s\n", opts.File)

			data, err := parser.Render(vars)
			if err != nil {
				return err
			}

			if !opts.Insert {
				fmt.Fprintln(os.Stdout, string(data))

				return nil
			}

			if opts.Backup {
				dest := fmt.Sprintf("%s_%s", opts.OutputFile, time.Now().Format("2006-01-02-150405"))

				if err := writer.CopyFile(opts.OutputFile, dest); err != nil {
					return err
				}

				fmt.Printf("created backup file %s\n", dest)
			}

			start, end, err := writer.SplitFile(opts.OutputFile, writer.Delimiter)
			if err != nil {
				return err
			}

			out := fmt.Sprintf("%s%s\n%s%s%s", start, writer.Delimiter, string(data), writer.Delimiter, end)

			//nolint: gosec
			if err := os.WriteFile(opts.OutputFile, []byte(out), 0o664); err != nil {
				return err
			}

			fmt.Printf("inserted vars into %s\n", opts.OutputFile)

			return nil
		},
	}

	cmd.AddCommand(versionCmd(version))

	cmd.Flags().StringVarP(&opts.File, "file", "f", opts.File, "path to the variables file")
	cmd.Flags().StringVarP(&opts.OutputFile, "output-file", "o", opts.OutputFile, "where to write the output to (required insert mode)")
	cmd.Flags().BoolVarP(&opts.Backup, "backup", "b", opts.Backup, "backup the output file before writing")
	cmd.Flags().BoolVarP(&opts.Insert, "insert", "i", opts.Insert, "insert mode, inserts the markdown table in the specified output file")
	cmd.Flags().BoolVarP(&opts.Markdown, "markdown", "m", opts.Markdown, "wether to create a markdown or html table")

	return cmd
}

// Execute invokes the command.
func Execute(version string) error {
	if err := newRootCmd(version).Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}
