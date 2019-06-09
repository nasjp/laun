package cmd

import (
	"github.com/NasSilverBullet/laun/internal/laun"
	"github.com/spf13/cobra"
)

func NewLaunCommand() *cobra.Command {
	cobra.OnInitialize()
	cmd := laun.Exec()
	return cmd
}
