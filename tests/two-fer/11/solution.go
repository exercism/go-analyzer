// Package twofer implements some sharing functionality.
package twofer

import "fmt"

// ShareWith shares with provided name.
func ShareWith(name string) string {
	if name != "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
