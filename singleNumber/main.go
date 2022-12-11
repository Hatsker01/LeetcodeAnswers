func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i, val := range m {
		if val == 1 {
			return i
		}
	}
	return 0
}	