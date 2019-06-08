package laun

import (
	"errors"
	"fmt"
	"strings"

	"github.com/NasSilverBullet/laun/pkg/apps"
	"github.com/NasSilverBullet/laun/pkg/config"
)

func open(args []string) error {
	if len(args) <= 0 {
		return errors.New(fmt.Sprintf("requires at least 1 arg(s), only received %d", len(args)))
	}
	c, err := config.New()
	if err != nil {
		return err
	}
	for _, arg := range args {
		lArg := strings.ToLower(arg)
		a := apps.New(lArg, c.AppPath)
		a.Target = target
		a, err = a.First()
		if err != nil {
			return err
		}
		err = a.Open()
		if err != nil {
			return err
		}
	}
	return err
}
