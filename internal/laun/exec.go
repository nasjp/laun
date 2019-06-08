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
		Long: `Plese use this simply.
You can do this for example with the following command:

   $ laun Docker
   $ laun docker            // You can use lower case
   $ laun doc               // You can use a part of apps name
   $ laun Docker Firefox    // You can luanch multiple applications at the same time
   $ laun atom -t README.md // You can launch with each file or directory
   $ laun -l                // You can check your applicatin list`,
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
