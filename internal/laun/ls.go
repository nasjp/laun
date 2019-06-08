package laun

import (
	"github.com/NasSilverBullet/laun/pkg/apps"
	"github.com/NasSilverBullet/laun/pkg/config"
)

func list() error {
	c, err := config.New()
	appsName, err := apps.All(c.AppPath)
	if err != nil {
		return err
	}
	apps.PrintList(appsName)
	return err
}
