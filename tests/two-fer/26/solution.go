// Package twofer contains ShareWith
package twofer

// ShareWith returns a message which depends on a name or not if no name is given
func ShareWith(name string) string {

	var res string
	res = name
	if len(name) <= 0 {
		res = "you"
	}

	return "One for " + res + ", one for me."
}
