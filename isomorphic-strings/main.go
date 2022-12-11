func isIsomorphic(s string, t string) bool {
	var a bool = false
	if s == t {
		return true
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && t[i] == t[j] {
				a = true
			} else if s[i] == s[j] && t[i] != t[j] {
				return false

			} else if s[i] != s[j] && t[i] == t[j] {
				return false

			} else {
				a = true
			}
		}
	}
	return a
}