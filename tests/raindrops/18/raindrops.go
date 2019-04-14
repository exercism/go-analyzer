package raindrops

import "strconv"

// Convert converts int number to string number. If the int number has 3 as a factor, output 'Pling', if 5 -> 'Plang', if 7 -> 'Plong'
func Convert(number int) string {
	contents := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	factors := [3]int{3, 5, 7}
	answer := ""
	for i := 0; i < 3; i++ {
		if number%factors[i] == 0 {
			answer += contents[factors[i]]
		}
	}
	if answer == "" {
		return strconv.Itoa(number)
	}
	return answer
}
