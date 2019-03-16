// Package twofer implements some sharing functionality.
package twofer

import "fmt"

// ShareWith shares with provided name.
func ShareWith(name string) string {
	if name == "" {
		return sharing("you")
	}
	return sharing(name)
}

func sharing(s string) string {
	return fmt.Sprintf("One for %s, one for me.", s)
}
