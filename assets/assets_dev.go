// +build !build

package assets

import "net/http"

// Patterns contains project assets.
var Patterns http.FileSystem = http.Dir("../patterns")
