package apps

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type App struct {
	Partial string
	Target string
	Name string
	Path string
}


//func (a *App) First() *App {
//
//}

func All(dir string) []App {
	return getApplications(dir)
}

func getApplications(dir string) (apps []App) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		a := &App{
			Name: file.Name(),
			Path :filepath.Join(dir, file.Name()),
		}

		//  exlusion dotfiles
		if strings.HasPrefix(a.Name, ".") {
			continue
		}

		// if dir
		if file.IsDir() && !strings.HasSuffix(a.Name, ".app") {
			apps = append(apps, getApplications(filepath.Join(dir, a.Name))...)
			continue
		}


	}

	return apps
}