// Package hamming provides the Distance function.
package hamming

import "fmt"

// Distance calculates the Hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("DNA strand lengths must be the same")
	}

	// Count how many nucleotides differ between the parameters
	distance := 0
	for i, na := range []byte(a) {
		if b[i] != na {
			distance++
		}
	}

	return distance, nil
}
