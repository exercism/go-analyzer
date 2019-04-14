package hamming

import (
	"errors"
)

var errDifferentLengths = errors.New("inputs are not the same length")

// Distance calculates the hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errDifferentLengths
	}

	var count int
	for i, val := range a {
		if val != []rune(b)[i] {
			count++
		}
	}
	return count, nil
}
