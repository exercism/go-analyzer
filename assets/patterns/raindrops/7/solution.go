package raindrops

import "strconv"

// Convert implements raindrop speech
func Convert(value int) string {
	var res string
	for i := 3; i <= 7; i += 2 {
		if value%i == 0 {
			switch i {
			case 3:
				res += "Pling"
			case 5:
				res += "Plang"
			case 7:
				res += "Plong"
			}
		}
	}

	if res == "" {
		return strconv.Itoa(value)
	}
	return res
}
