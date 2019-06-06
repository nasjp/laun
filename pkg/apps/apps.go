package apps

import (
	"fmt"
	"io/ioutil"
	"log"
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

func New(partial string, target string, path string) *App {
	return &App{
		Partial : partial,
		Target : target,
		Path : path,
	}
}




type app interface {
	First() *App
}

func (a *App) First() *App {
	files, err := ioutil.ReadDir(a.Path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if name:= file.Name(); strings.Contains(name, a.Partial) {
			a.Name = name
			break
		}
	}
	return a
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
