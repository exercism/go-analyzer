package hamming

import (
	"errors"
)

// Distance calculation hamming distance for two strings
func Distance(a, b string) (int, error) {
	if len([]rune(a)) != len([]rune(b)) {
		return 0, errors.New("not equal strings")
	}
	var diff int
	for i, r := range []rune(a) {
		if r != rune(b[i]) {
			diff++
		}
	}
	return diff, nil
}
