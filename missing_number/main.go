package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(a))
}

func firstMissingPositive(nums []int) int {
		sort.Ints(nums)
		var t, v bool = false, false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == 1 {
				v = true
			}
			if nums[i] != nums[i+1]-1 && nums[i] > 0 && v {

				return nums[i] + 1
			}
			if nums[i] == len(nums) {
				t = true
			}

		}
		if !v {
			return 1
		}
		if !t {
			return len(nums)
		}

		return len(nums)
}
