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
		apps := All(dir);
		for _, a := range apps {
			if !strings.HasSuffix(a.Path, ".app") {
				t.Errorf("All(%s) = %s", dir, a.Name)
			}
		}
	}
}