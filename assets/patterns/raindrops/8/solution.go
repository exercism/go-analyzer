package raindrops

import (
	"strconv"
	"strings"
)

// Convert int to raindrop speak
func Convert(num int) string {
	res := strings.Builder{}
	if num%3 == 0 {
		res.WriteString("Pling")
	}
	if num%5 == 0 {
		res.WriteString("Plang")
	}
	if num%7 == 0 {
		res.WriteString("Plong")
	}
	if res.Len() == 0 {
		res.WriteString(strconv.Itoa(num))
	}
	return res.String()
}
