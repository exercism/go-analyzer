// Package twofer provides directions for how to split items.
package twofer

import "fmt"

// ShareWith describes how many for whom.
func ShareWith(name string) string {
	if len(name) > 0 {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
	return "One for you, one for me."
}
