package laun

import (
	"github.com/NasSilverBullet/laun/pkg/apps"
	"github.com/NasSilverBullet/laun/pkg/config"
)

func list() error {
	c := config.New()
	apps.PrintList(apps.All(c.AppPath))
	return nil
}
