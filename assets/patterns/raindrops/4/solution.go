package raindrops

import "strconv"

var speech = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert implements raindrop speech
func Convert(value int) string {
	var res string
	for i := 3; i <= 7; i += 2 {
		if value%i == 0 {
			res += speech[i]
		}
	}

	if res == "" {
		res = strconv.Itoa(value)
	}
	return res
}
