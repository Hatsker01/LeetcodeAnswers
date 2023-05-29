package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[count] != nums[i] {
			count++
			nums[count] = nums[i]
		}
	}
	return count + 1
}
