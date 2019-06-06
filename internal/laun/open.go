package laun

import (
	"github.com/YukihiroTaniguchi/laun/pkg/apps"
	"github.com/YukihiroTaniguchi/laun/pkg/config"
	"github.com/spf13/cobra"
)

var t string

func open() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "open application",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := config.New()
			for _, arg := range args {
				a := apps.New(arg, c.AppPath)
				a.Target = t
				a = a.First()
				a.Open()
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&t, "target", "t", "", "select target file or directory")
	return cmd
}