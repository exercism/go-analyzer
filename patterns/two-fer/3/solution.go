// Package twofer implements some sharing functionality.
package twofer

import "fmt"

// ShareWith shares with provided name.
func ShareWith(name string) string {
	if name == "" {
		return share("you")
	}
	return share(name)
}

func share(name string) string {
	return fmt.Sprintf("One for %s, one for me.", name)
}
