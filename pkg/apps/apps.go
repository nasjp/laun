package apps

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type App struct {
	Partial string
	Target string
	Name string
	Path string
}

var suffix = ".app"

//func (a *App) First() *App {
//
//}

func All(dir string) (apps []App) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}

		if !strings.HasSuffix(name, suffix) {
			continue
		}

		apps = append(apps, App{
			Name: strings.TrimSuffix(name, suffix),
			Path: filepath.Join(dir, name),
		})
	}

	return apps
}