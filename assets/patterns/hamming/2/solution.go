package hamming

import "errors"

// Distance returns the hamming distance of given strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings have different length")
	}

	diff := 0
	bRunes := []rune(b)
	for i, r := range a {
		if r != bRunes[i] {
			diff++
		}
	}
	return diff, nil
}
