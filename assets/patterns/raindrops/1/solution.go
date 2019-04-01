package raindrops

import "strconv"

// Convert implements raindrop speech
func Convert(i int) string {
	var res string
	if i%3 == 0 {
		res += "Pling"
	}
	if i%5 == 0 {
		res += "Plang"
	}
	if i%7 == 0 {
		res += "Plong"
	}

	if res == "" {
		res = strconv.Itoa(i)
	}

	return res
}
