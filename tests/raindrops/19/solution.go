// Package raindrops implements raindrop functionality
package raindrops

import "strconv"

var speech = []struct {
	v int
	s string
}{
	{v: 3, s: "Pling"},
	{v: 5, s: "Plang"},
	{v: 7, s: "Plong"},
}

// Convert int to raindrop speech
func Convert(n int) string {
	var out string
	for _, def := range speech {
		if n%def.v == 0 {
			out += def.s
		}
	}

	if out == "" {
		return strconv.Itoa(n)
	}
	return out
}
