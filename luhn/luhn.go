package luhn

import (
	"fmt"
	"log"
	"regexp"
)

func Valid(input string) bool {
	status := false
	if len(input) > 1 {
		stringResult := ""
		fmt.Println(input, "new")
		inputNew := getOnlyNumeric(input)
		fmt.Println(inputNew, "inputNew")
		for i := 0; i < len(inputNew); i++ {
			if (i+1)%2 != 0 {
				doubling := int(inputNew[i]) * 2
				if doubling > 9 {
					doubling = doubling - 9
				}
				fmt.Println(inputNew[i], doubling, "doubling")
				stringResult += fmt.Sprint(doubling)
			} else {
				stringResult += string(inputNew[i])
			}
		}
		sum := 0

		for i := 0; i < len(inputNew); i++ {
			sum += int(inputNew[i])
		}

		if sum%10 == 0 {
			status = true
		}

	}

	return status
}

func getOnlyNumeric(s string) string {

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(s, "")

}
