package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "version",
		Short:         "print ansdoc version",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "ansdoc %s\n", version)

			return nil
		},
	}

	return cmd
}
