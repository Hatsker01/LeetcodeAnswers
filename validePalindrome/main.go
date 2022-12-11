func isPalindrome(s string) bool {
	var str string = ""
	// var t bool=false
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			str += string(s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			str += string(s[i])
		} else if s[i] >= '0' && s[i] <= '9' {
			str += string(s[i])
		}
	}
	str = strings.ToLower(str)

	for i := 0; i < len(str); i++ {
		if string(str[i]) != string(str[len(str)-i-1]) {

			return false
		}
	}
	return true
}