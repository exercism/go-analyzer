// Package twofer implements some sharing functionality.
package twofer

import "fmt"

// SomeName shares with provided name.
func SomeName(name string) string {
	n := name
	if name == "" {
		n = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", n)
}
