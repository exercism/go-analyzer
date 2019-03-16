// Package twofer implements some sharing functionality.
package twofer

import "fmt"

// ShareWith shares with provided name.
func ShareWith(name string) string {
	n := name
	if name == "" {
		n = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", n)
}
