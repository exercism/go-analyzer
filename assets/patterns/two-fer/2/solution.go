// Package twofer contains sharing algorithms.
package twofer

import "fmt"

// ShareWith shares with given name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
