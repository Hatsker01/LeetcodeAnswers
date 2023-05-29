package main

import "sort"

func main(){

}

func thirdMax(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-4]
}