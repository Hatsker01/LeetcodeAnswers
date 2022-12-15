func sortSentence(s string) string {
	m := make(map[int]string)
	var ints []int
	var t string = ""
	var f int = 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			j, _ := strconv.Atoi(string(s[i]))
			m[j] = string(s[f:i])

			f = i + 2
			ints = append(ints, j)
		} else if string(s[i]) == " " {
			continue
		}

	}

	sort.Ints(ints)

	for i := 0; i < len(ints); i++ {
		if i == len(ints)-1 {
			t += m[ints[i]]
		} else {
			t += m[ints[i]] + " "

		}
	}
	return t
}