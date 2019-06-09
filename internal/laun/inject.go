package laun

import "github.com/spf13/cobra"

func injectFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringVarP(&target, "target", "t", "", "select target file or directory")
	cmd.Flags().BoolVarP(&ls, "list", "l", false, "show application list")
	return cmd
}
