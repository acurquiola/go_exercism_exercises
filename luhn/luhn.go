package luhn

import (
	"strconv"
	"strings"
)

func Valid(input string) bool {
	status := false
	input = strings.ReplaceAll(input, " ", "")

	if len(input) > 1 {
		sum := 0
		auxMod := len(input) % 2

		for index, char := range strings.Split(input, "") {
			digit, err := strconv.Atoi(char)
			if err != nil {
				return false
			}

			if auxMod == (index % 2) {
				digit = digit * 2
			}

			if digit > 9 {
				digit = digit - 9
			}

			sum += digit
		}

		if (sum % 10) == 0 {
			status = true
		}
	}

	return status
}
