// Package hamming implements hamming distance
package hamming

import (
	"errors"
)

// Distance calculates the hamming distance between two string
func Distance(a, b string) (int, error) {
	var (
		runesA = []rune(a)
		runesB = []rune(b)
	)
	if len(runesA) != len(runesB) {
		return 0, errors.New("mismatch length strings not accepted")
	}

	hammingDistance := 0
	for i, char := range runesA {
		if char != runesB[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil

}
