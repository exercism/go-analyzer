// Package twofer provides directions for how to split items.
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith describes how many for whom.
func ShareWith(name string) string {
	if strings.Compare(name, "") == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
