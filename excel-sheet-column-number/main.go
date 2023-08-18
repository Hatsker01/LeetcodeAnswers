package main

import "fmt"

func Valid(s string) bool {
	if len(s) == 1 {
		return false
	} else if s == "" {
		return true
	} else if (s[0] == '{' && s[1] == '}') || (s[0] == '[' && s[1] == ']') || (s[0] == '(' && s[1] == ')') {
		s = s[2:]
		return Valid(s)
	} else if (s[0] == '{' && s[len(s)-1] == '}') || (s[0] == '[' && s[len(s)-1] == ']') || (s[0] == '(' && s[len(s)-1] == ')') {
		s = s[1 : len(s)-1]
		return Valid(s)
	}
	return false
}
func main() {
	fmt.Println(Valid("(([]){})"))

	// fmt.Println(convertToTitle(28))
}
