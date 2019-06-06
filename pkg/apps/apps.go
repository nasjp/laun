package apps

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	suffix = ".app"
	format = "\x1b[32m%s\x1b[0m" // 32m => green
)

type App struct {
	Partial string
	Target string
	Path string
	Name string
}

func New(partisal string, path string) *App {
	return &App{
		Partial:partisal,
		Path : path,
	}
}

type app interface {
	First() *App
	Open(*App)
}

func (a *App) First() *App {
	files, err := ioutil.ReadDir(a.Path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if name:= strings.ToLower(file.Name()); strings.Contains(name, a.Partial) {
			a.Name = name
			break
		}
	}
	return a
}

func (a *App) Open() {
	var err error
	if len(a.Target) > 0 {
		err = exec.Command("open", a.Target, "-a", a.Name).Start()
	}else{
		err = exec.Command("open", "-a", a.Name).Start()
	}
	if err != nil {
		log.Fatal(err)
	}
}

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

func PrintList(apps []App) {
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
