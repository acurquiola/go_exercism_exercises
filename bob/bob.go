package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	//By default Bob says
	bobAnswer := "Whatever."

	is_only_numbers := CheckOnlyNumber(remark)
	remark = string(strings.TrimSpace(remark))
	len_remark := len(remark)

	if len_remark > 0 && !isSilence(remark) {
		if strings.Contains(remark, "?") && (strings.ToUpper(string(remark)) == remark && !is_only_numbers) {
			bobAnswer = "Calm down, I know what I'm doing!"
			//If remark is a question
		} else if remark[len(remark)-1:] == "?" {
			bobAnswer = "Sure."
			//If remark is yell sentence
		} else if strings.ToUpper(string(remark)) == remark && !is_only_numbers {
			bobAnswer = "Whoa, chill out!"
		}
	} else {
		//If is a silence
		bobAnswer = "Fine. Be that way!"
	}

	return bobAnswer
}

func CheckOnlyNumber(str string) bool {
	letter_counter := 0
	flag := false
	for _, charVariable := range str {

		if unicode.IsLetter(charVariable) {
			letter_counter++
		}
	}

	if letter_counter == 0 {
		flag = true
	}

	return flag
}

func isSilence(str string) bool {
	flag := true
	for _, charVariable := range str {
		if !unicode.IsSpace(charVariable) {
			flag = false
		}
	}

	return flag
}
