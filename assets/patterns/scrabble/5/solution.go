package scrabble

import "unicode"

const (
	asciiA = 'a'
	asciiZ = 'z'
)

var scores = []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

// Score implements scrabble score
func Score(s string) int {
	var score int
	for _, r := range s {
		r = unicode.ToLower(r)
		if r < asciiA || asciiZ < r {
			continue
		}
		score += scores[r-asciiA]
	}

	return score
}
