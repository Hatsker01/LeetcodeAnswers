package main

import "sort"

func findDuplicate(nums []int) []int {
	sort.Ints(nums)
	m := make(map[int]int)
	var n []int
	
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	if nums[0]!=1{
		for k, v := range m {
		
			if v > 1 {
				n = append(n, k-1)
			}
		}
	}else{
		for k, v := range m {
		
			if v > 1 {
				n = append(n, k+1)
			}
		}
	}
	
	return n
}
