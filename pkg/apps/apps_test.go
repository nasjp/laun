package apps

import (
	"strings"
	"testing"
)


func TestAll(t *testing.T) {
	var dirs = []string{
		"/Applications",
	}
	a := NewApp()
	for _, dir := range dirs {
		apps := a.All(dir);
		for _, app := range apps {
			if !strings.HasSuffix(app.Path, ".app") {
				t.Errorf("All(%s) = %s", dir, app.Name)
			}
		}
	}
}

func TestPrintList(t *testing.T) {
	apps := []App{
		{Name: "a.app"},
	}
		a := NewApp()
		a.PrintList(apps)
}