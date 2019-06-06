package laun

import (
	"github.com/YukihiroTaniguchi/laun/pkg/apps"
	"github.com/YukihiroTaniguchi/laun/pkg/config"
	"github.com/spf13/cobra"
)

var (
	format = "\x1b[32m%s\x1b[0m" // 32m => green
)


func ls() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "show application list",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := config.New()
			a := apps.NewApp()
			a.PrintList(a.All(c.AppPath))
			return nil
		},
	}
	return cmd
}
