package acronym

import (
	"log"
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	acronym := ""
	if len(s) > 0 {
		s = strings.ReplaceAll(s, "-", " ")
		s_split := Explode(" ", s)
		for _, v := range s_split {
			v = getOnlyAlphabetic(v)
			if v != "" {
				first_letter := v[0:1]
				acronym += first_letter
			}
		}
	}
	return strings.ToUpper(acronym)
}

func getOnlyAlphabetic(s string) string {
	if len(s) > 0 {
		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			log.Fatal(err)
		}

		processedString := reg.ReplaceAllString(s, "")

		return processedString
	}

	return ""
}

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}
