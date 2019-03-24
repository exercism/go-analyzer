// +build !build

package assets

import "net/http"

// Patterns contains good patterns per exercise for pattern matching.
var Patterns http.FileSystem = http.Dir("assets/patterns")
