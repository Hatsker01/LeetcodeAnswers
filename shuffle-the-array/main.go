package main

func main() {

}

func shuffle(nums []int, n int) []int {
	var arra []int = []int{}
	for i := 0; i < n; i++ {
		arra = append(arra, nums[i], nums[n+i])
	}
	return arra
}

