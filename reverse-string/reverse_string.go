package reverse

func Reverse(input string) string {
	reverseWord := ""
	if len(input) > 0 {
		for _, i := range input {
			reverseWord = string(i) + reverseWord
		}
	}
	return reverseWord
}
