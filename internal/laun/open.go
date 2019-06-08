package laun

import (
	"errors"
	"fmt"
	"strings"

	"github.com/YukihiroTaniguchi/laun/pkg/apps"
	"github.com/YukihiroTaniguchi/laun/pkg/config"
)

func open(args []string) error {
	if len(args) <= 0 {
		return errors.New(fmt.Sprintf("requires at least 1 arg(s), only received %d", len(args)))
	}
	c := config.New()
	for _, arg := range args {
		lArg := strings.ToLower(arg)
		a := apps.New(lArg, c.AppPath)
		a.Target = target
		a = a.First()
		a.Open()
	}
	return nil
}
