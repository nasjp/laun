package apps

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	var dirs = []string{
		"/Applications",
	}
	for _, dir := range dirs {
		apps, err := All(dir)
		if err != nil {
			t.Fatal(err)
		}
		for _, app := range apps {
			if !strings.HasSuffix(app.Path, ".app") {
				t.Errorf("All(%s) = %s", dir, app.Name)
			}
		}
	}
}
