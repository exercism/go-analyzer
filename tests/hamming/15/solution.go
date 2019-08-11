package hamming

import "fmt"

var a, b string = "GTCAGTCA", "GTAACTGT"

func Distance(a, b string) (int, error) {
	hammingDistance := 0
	error := fmt.Errorf("cannot compare; strings are not the same length")
	if len(a) == len(b) {
		fmt.Printf("a length: %v, b length: %v", len(a), len(b))
		if len(a) == 0 {
			return hammingDistance, nil
		} else {
			for aIndex, aValue := range a {
				for bIndex, bValue := range b {
					if aIndex == bIndex && aValue != bValue {
						fmt.Printf("\na value: %v, b value: %v", aValue, bValue)
						hammingDistance += 1
					}
				}
			}
		}
	} else {
		return hammingDistance, error
	}
	return hammingDistance, nil
}
