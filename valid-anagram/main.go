package main

func main() {

}

func isAnagram(s string, t string) bool {
	m := make(map[string]int)

	n := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}
	for i := 0; i < len(t); i++ {
		n[string(t[i])]++
	}
	for c, v := range m {
		if n[c] != v {
			return false
		}
	}
	for c, v := range n {
		if m[c] != v {
			return false
		}
	}
	return true
}
