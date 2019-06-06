package laun

import "github.com/spf13/cobra"

func Inject(cmd *cobra.Command) *cobra.Command {
	cmd.AddCommand(ls())
	cmd.AddCommand(open())
	return cmd
}