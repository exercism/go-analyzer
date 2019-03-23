// Package twofer provides a routine for two-for-one strings
package twofer

import "strings"

// ShareWith should return "One for X, one for me.", where X is either a name or "you"
func ShareWith(name string) string {
	var str strings.Builder

	str.WriteString("One for ")

	if name != "" {
		str.WriteString(name)
	} else {
		str.WriteString("you")
	}

	str.WriteString(", one for me.")

	return str.String()
}
