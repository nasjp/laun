package laun

import "github.com/spf13/cobra"

func Inject(cmd *cobra.Command) *cobra.Command {
	ls := ls()
	cmd.AddCommand(ls)
	return cmd
}