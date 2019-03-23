// Package twofer shares things.
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith shares something with someone else. It defaults to "you".
func ShareWith(name string) string {
	if strings.TrimSpace(name) == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
