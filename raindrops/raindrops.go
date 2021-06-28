package raindrops

import (
	"fmt"
	"math"
	"strings"
)

func Convert(number int) string {
	factors := [3]float64{3.00, 5.00, 7.00}
	text := [3]string{"Pling", "Plang", "Plong"}
	result := ""
	mod := 0.00
	float_number := float64(number)
	string_number := fmt.Sprint(number)

	for i := 0; i < len(factors); i++ {
		mod = math.Mod(float_number, factors[i])
		if result == "" && mod > 0 {
			result = string_number
		} else {
			if mod == 0 {
				result += text[i]
				result = strings.ReplaceAll(result, string_number, "")
			}
		}
	}

	return result
}
