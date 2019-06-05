package config

import "runtime"

const (
	M   = "darwin"
	L   = "linux"
	W = "windows"
)

type Config struct {
	OS              string
	AppPath string
}

func New() *Config {
	var path string

	switch runtime.GOOS {
	case M:
		path = "/Applications"
	case L:
		path = "/opt"
	case W:
		path = "C:¥¥Program Files"
	}

	return &Config{
		OS : runtime.GOOS,
		AppPath : path,
	}
}