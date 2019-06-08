package config

import (
	"errors"
	"fmt"
	"runtime"
)

const (
	M = "darwin"
	L = "linux"
	//W = "windows"
)

type Config struct {
	OS      string
	AppPath string
}

func New() (*Config, error) {
	var (
		path string
		err  error
	)

	switch runtime.GOOS {
	case M:
		path = "/Applications"
	case L:
		path = "/opt"
	//case W:
	//	path = "C:¥¥Program Files"
	default:
		err = errors.New(fmt.Sprintf("I'm sorry, %s os cant't use laun", runtime.GOOS))
	}
	c := &Config{
		OS:      runtime.GOOS,
		AppPath: path,
	}
	return c, err
}
