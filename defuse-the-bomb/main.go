func decrypt(code []int, k int) []int {
	var end []int
	var t int = 0
	var y int = 0
	if k == 0 {
		for i := 0; i < len(code); i++ {
			end = append(end, 0)
		}
		return end
	} else if k > 0 {
		for i := 1; i <= len(code); i++ {
			for j := i; y != k; j++ {
				if j > len(code)-1 {
					t += code[j-len(code)]

					y++
				} else {
					t += code[j]

					y++
				}
			}
			end = append(end, t)
			y = 0
			t = 0
		}
	} else {
		for i := len(code) - 1; i >= 0; i-- {
			for j := len(code) - i - 2; y != -k; j-- {
				if j < 0 {
					t += code[j+len(code)]

					y++
				} else {
					t += code[j]

					y++
				}
			}
			end = append(end, t)
			y = 0
			t = 0
		}
	}

	return end
}