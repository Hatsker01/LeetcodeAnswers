package main

func moveZeroes(nums []int) {

	var a []int
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			a = append(a, nums[i])
		}

	}
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		a = append(a, 0)
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = a[i]
	}
}
