package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3,2,2,3},3))

}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			
			nums[count] = nums[i]
			// fmt.Println("1-",nums[count])
			
		} else{
			nums[count]=nums[i]
			// fmt.Println("2-",nums[count])
			count++
		}
		
	}
	fmt.Println(nums)
	return count 
}
