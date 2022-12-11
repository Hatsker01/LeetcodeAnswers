func romanToInt(s string) int {
	sum := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		val := m[s[i]]

		sign := 1
		if last > val {
			sign = -1
		}
		sum += sign * val
		last = val
	}
	return sum

}