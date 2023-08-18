package main

func main() {

}

func numIdenticalPairs(nums []int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[i] == nums[k] {
				count++
			}
		}
	}
	return count
}
