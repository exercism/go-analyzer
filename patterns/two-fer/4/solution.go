// Package twofer implements some sharing functionality.
package twofer

// ShareWith shares with provided name.
func ShareWith(name string) string {
	if name == "" {
		return share("you")
	}
	return share(name)
}

func share(name string) string {
	return "One for " + name + ", one for me."
}
