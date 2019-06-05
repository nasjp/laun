package laun

import (
	"fmt"
	"github.com/YukihiroTaniguchi/laun/pkg/apps"
	"github.com/YukihiroTaniguchi/laun/pkg/config"
	"github.com/spf13/cobra"
)

func ls() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "show application list",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := config.New()

			apps := apps.All(c.AppPath)
			fmt.Println(apps)
			return nil
		},
	}
	return cmd
}