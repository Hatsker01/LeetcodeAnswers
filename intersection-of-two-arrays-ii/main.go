package main

import "fmt"

func main() {
	a := []int{1, 2, 2, 1}
	b := []int{2, 2}
	fmt.Println(intersect(a, b))
}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	n := make(map[int]int)
	var a []int
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		n[nums2[i]]++
	}
	
	if len(nums2 )<len(nums1) {
		for i, v := range n {
			if m[i] > 0 {
				if v>m[i]{
					v=m[i]
				}
				for k := 0; k < v; k++ {
					a = append(a, i)
				}
			}
		}
	}else{
		for i, v := range m {
			if n[i] > 0 {
				if v>n[i]{
					v=n[i]
				}
				for k := 0; k < v; k++ {
					a = append(a, i)
				}
			}
		}
	}
	return a

}
