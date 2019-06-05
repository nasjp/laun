package cmd

import (
	"github.com/YukihiroTaniguchi/laun/internal/apps/laun"
	"github.com/spf13/cobra"
)



func NewLaunCommand(wd string) *cobra.Command {
	cobra.OnInitialize()
	cmd := laun.Exec()
	cmd = laun.Inject(cmd)
	return cmd
}

