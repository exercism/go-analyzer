// Package hamming promises to improve your productivity
// 100x by computing your hamming distances.
package hamming

import (
	"errors"
	"strings"
)

// Distance computes the hamming distance of two strings.
func Distance(a, b string) (int, error) {
	dist, err := 0, error(nil)
	if len(a) != len(b) {
		err = errors.New("strings must be same length")
	} else {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
		for i, c := range a {
			if string(c) != string(b[i]) {
				dist++
			}
		}
	}
	return dist, err
}
