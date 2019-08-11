// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith shares a piece with you
func ShareWith(name string) string {
	you := name
	if you == "" {
		you = "you"
	}
	return "One for " + you + ", one for me."
}
