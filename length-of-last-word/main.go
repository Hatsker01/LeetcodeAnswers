func lengthOfLastWord(s string) int {
	var length int
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			continue
		} else if i >= 1 && string(s[i-1]) == " " && string(s[i]) != " " {
			length = 0
			length++
		} else if string(s[0]) != " " {
			length++
		} else if i >= 1 && string(s[i]) != " " && string(s[i-1]) != " " {

			length++

		}
		// fmt.Println(i," ",string(s[i])," ",length)
	}
	return length
}