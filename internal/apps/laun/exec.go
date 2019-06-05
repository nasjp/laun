package laun

import (
	"github.com/spf13/cobra"
)

func Exec() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "laun",
		Short: "laun is simple application launcher",
		//RunE: func(cmd *cobra.Command, args []string) error {
		//	fmt.Println("execute laun")
		//	return nil
		//},
	}
	return cmd
}