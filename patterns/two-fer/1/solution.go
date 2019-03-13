// Package twofer contains sharing algorithms.
package twofer

// ShareWith shares with given name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
