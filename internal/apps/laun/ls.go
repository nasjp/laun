package laun

import (
	"fmt"
	"github.com/spf13/cobra"
)

func ls() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "show application list",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("apps")
			return nil
		},
	}
	return cmd
}