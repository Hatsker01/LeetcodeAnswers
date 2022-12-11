func titleToNumber(columnTitle string) int {
	var sum int = 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		if i == len(columnTitle)-1 {
			sum += int(columnTitle[i] - 64)
		} else {
			sum += int(columnTitle[i]-64) * int(math.Pow(26, float64(len(columnTitle)-i-1)))
		}
	}
	return sum
}