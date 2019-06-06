package apps

import (
	"strings"
	"testing"
)

func TestGetApplications(t *testing.T) {
	var dirs = []string{
		"/Applications",
	}
	for _, dir := range dirs {
		apps := getApplications(dir);
		for _, a := range apps {
			if !strings.HasSuffix(a.Path, ".app") {
				t.Errorf("getApplications(%s) = %s", dir, a.Name)
			}
		}
	}
}