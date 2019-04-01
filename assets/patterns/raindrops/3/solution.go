package raindrops

import "strconv"

var speech = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert implements raindrop speech
func Convert(i int) string {
	var res string
	for _, mod := range []int{3, 5, 7} {
		if i%mod == 0 {
			res += speech[mod]
		}
	}

	if res == "" {
		res = strconv.Itoa(i)
	}
	return res
}
