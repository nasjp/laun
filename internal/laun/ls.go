package laun

import (
	"github.com/YukihiroTaniguchi/laun/pkg/apps"
	"github.com/YukihiroTaniguchi/laun/pkg/config"
)

func list() error {
	c := config.New()
	apps.PrintList(apps.All(c.AppPath))
	return nil
}
