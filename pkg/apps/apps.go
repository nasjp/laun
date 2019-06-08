package apps

import (
	"errors"
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
	Target  string
	Path    string
	Name    string
}

func New(partisal string, path string) *App {
	return &App{
		Partial: partisal,
		Path:    path,
	}
}

type app interface {
	First() *App
	Open()
}

func (a *App) First() (*App, error) {
	files, err := ioutil.ReadDir(a.Path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if name := strings.ToLower(file.Name()); strings.Contains(name, a.Partial) {
			a.Name = name
			break
		}
	}
	if a.Name == "" {
		err = errors.New(fmt.Sprintf("Unable to find application that name contains '%s'", a.Partial))
		return nil, err
	}
	return a, err
}

func (a *App) Open() error {
	// TODO: エラー処理を修正する
	var err error
	if len(a.Target) > 0 {
		err = exec.Command("open", a.Target, "-a", a.Name).Start()
	} else {
		err = exec.Command("open", "-a", a.Name).Start()
	}
	return err
}

func All(dir string) ([]App, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var apps []App
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
	if len(apps) <= 0 {
		err := errors.New(fmt.Sprintf("No application"))
		return apps, err
	}
	return apps, nil
}

func PrintList(apps []App) {
	var list, lineFeed string
	for _, app := range apps {
		list = list + func() string {
			line := lineFeed + app.Name
			lineFeed = "\n"
			return line
		}()
	}
	fmt.Printf(format+"\n", list)
}
