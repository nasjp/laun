package apps

import (
	"fmt"
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


var (
	suffix = ".app"
	format = "\x1b[32m%s\x1b[0m" // 32m => green
)

type AppInterface struct {
}

type appInterface interface {
	All(string) []App
	PrintList([]App)

}

func NewApp() *AppInterface {
	return &AppInterface{}
}


//func (a *App) First() *App {
//
//}

func (*AppInterface) All(dir string) (apps []App) {
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

func (*AppInterface) PrintList(apps []App) {
	var list, lineFeed string
	for _, app := range apps {
		list = list + func() string{
			line := lineFeed + app.Name
			lineFeed = "\n"
			return line
		}()
	}
	fmt.Printf(format + "\n", list)
}
