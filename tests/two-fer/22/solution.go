// Package twofer has a share function.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns a string following the two-fer rules.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
