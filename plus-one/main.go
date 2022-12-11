func plusOne(digits []int) []int {
	var number string = ""
	var arra []int
	for _, i := range digits {
		number += strconv.Itoa(i)
	}
	var add int = 0
	for i := 0; i < len(number); i++ {
		t, _ := strconv.Atoi(string(number[len(number)-i-1]))

		if i == 0 {
			if t+1 >= 10 {
				arra = append(arra, (t+1)%10)
				add = 1
			} else {
				arra = append(arra, t+1)
				add = 0
			}

		} else if add > 0 {
			if t+1 >= 10 {
				arra = append(arra, (t+1)%10)
				add = 1
			} else {
				arra = append(arra, t+1)
				add = 0
			}
		} else {
			arra = append(arra, t)
			add = 0
		}
	}
	if add == 1 {
		arra = append(arra, 1)
	}
	var ans []int
	for i := 0; i < len(arra); i++ {
		ans = append(ans, arra[len(arra)-i-1])
	}

	return ans
}