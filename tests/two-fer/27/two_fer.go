// Package twofer has methods for the common "One for you, one for me" saying.
package twofer

// ShareWith returns a string of the famous 'One for you, one for me' with you replaced by `name`.
func ShareWith(name string) string {
	prefix := "One for "
	suffix := ", one for me."

	if name == "" {
		name = "you"
	}
	return prefix + name + suffix
}
