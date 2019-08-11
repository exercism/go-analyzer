package raindrops

import (
	"math"
	"sort"
	"strconv"
)

// Convert a number to a string, the contents of which depend on the number's factors.
// If the number has 3 as a factor, output 'Pling'.
// If the number has 5 as a factor, output 'Plang'.
// If the number has 7 as a factor, output 'Plong'.
// If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
func Convert(input int) string {

	i := 2
	var output string
	factors := make(map[int]struct{}, 0)

	if input < 3 {
		return strconv.Itoa(input)
	}
	if input == 3 || input == 5 || input == 7 {
		return getPlingPlangOrPong(input)
	}

	check := math.Ceil(math.Sqrt(float64(input)))

	for {

		if input%i == 0 {
			factors[i] = struct{}{}
			if i != (input / i) {
				factors[input/i] = struct{}{}
			}
		}

		if float64(i) > check {
			break
		}
		i++
	}

	j := make([]int, 0)

	for key := range factors {
		j = append(j, key)
	}

	for _, number := range j {
		output += getPlingPlangOrPong(number)
	}

	sort.Ints(j)

	if output == "" {
		output = strconv.Itoa(input)
	}

	return output
}

func getPlingPlangOrPong(input int) string {
	switch input {
	case 3:
		return "Pling"
	case 5:
		return "Plang"
	case 7:
		return "Plong"
	}

	return ""
}
