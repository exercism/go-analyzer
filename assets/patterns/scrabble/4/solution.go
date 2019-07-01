package scrabble

import (
	"unicode"
)

// Score calculates scrabble score for the given word.
func Score(word string) int {
	var sum int
	for _, r := range word {
		sum = getScore(r)
	}
	return sum
}

func getScore(r rune) int {
	switch unicode.ToUpper(r) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}
	return 0
}
