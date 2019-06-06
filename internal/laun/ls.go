package laun

import (
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
			apps.PrintList(apps.All(c.AppPath))
			return nil
		},
	}
	return cmd
}
