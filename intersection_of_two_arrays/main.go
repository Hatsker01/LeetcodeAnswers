package main

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var n []int
	for i := 0; i < len(nums1); i++ {
		for k := 0; k < len(nums2); k++ {
			if nums1[i] == nums2[k] {
				m[nums1[i]]++
			}
		}
	}
	for i, v := range m {
		if v >= 1 {
			n = append(n, i)
		}
	}
	return n
}
