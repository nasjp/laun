package laun

import (
	"github.com/YukihiroTaniguchi/laun/pkg/apps"
	"github.com/YukihiroTaniguchi/laun/pkg/config"
)

func open(args []string) error {
	c := config.New()

	for _, arg := range args {
		a := apps.New(arg, c.AppPath)
		a.Target = target
		a = a.First()
		a.Open()
	}
	return nil
}
