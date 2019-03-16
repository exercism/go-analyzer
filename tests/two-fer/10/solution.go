/*
Package twofer or 2-fer is short for two for one.
One for you and one for me.
*/
package twofer

import "fmt"

// ShareWith shares with provided name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
