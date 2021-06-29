package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	isBoolean := true

	for _, w := range word {
		if unicode.IsLetter(w) {
			if strings.Count(strings.ToUpper(word), strings.ToUpper(string(w))) > 1 {
				isBoolean = false
			}
		}
	}

	return isBoolean
}
