// +build !build

package assets

import (
	"net/http"
	"path"
	"runtime"
)

// Patterns contains good patterns per exercise for pattern matching.
var Patterns http.FileSystem = http.Dir("assets/patterns")

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return
	}
	Patterns = http.Dir(path.Join(path.Dir(filename), "patterns"))
}
