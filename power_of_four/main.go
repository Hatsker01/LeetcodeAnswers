func isPowerOfFour(n int) bool {
	var number int = 0
	for i := 0; number < n; i++ {
		if int(math.Pow(4, float64(i))) == n {
			return true
		}
		number = int(math.Pow(4, float64(i)))
	}
	return false
}