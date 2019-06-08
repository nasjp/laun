package laun

import (
	"github.com/spf13/cobra"
)

var (
	target string
	ls     bool
)

func Exec() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "laun",
		Short: "laun is simple application launcher",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if ls {
				err := list()
				return err
			}
			err := open(args)
			return err
		},
	}
	return cmd
}
