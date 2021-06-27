package twofer

import "strings"

func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return strings.ReplaceAll("One for _NAME_, one for me.", "_NAME_", name)
}
