// Package twofer implements some sharing functionality.
package twofer

import "fmt"

// ShareWith shares with provided name.
func ShareWith() string {
	n := "Bob"
	return fmt.Sprintf("One for %s, one for me.", n)
}
