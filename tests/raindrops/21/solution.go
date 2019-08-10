package raindrops

import (
	"strconv"
)

// Convert implements raindrop speech
func Convert(i int) string {
	var res string
	if hasFactor(i, 3) {
		res += "Pling"
	}
	if hasFactor(i, 5) {
		res += "Plang"
	}
	if hasFactor(i, 7) {
		res += "Plong"
	}

	if res == "" {
		return strconv.Itoa(i)
	}

	return res
}

func hasFactor(n, f int) bool {
	return n%f == 0
}
