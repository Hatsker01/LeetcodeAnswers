func searchInsert(nums []int, target int) int {
	var place int = 0
	for i, val := range nums {
		if val > target {
			place = i
			return i
		}
		if val == target {

			place = i
			return i

		} else {
			place = len(nums)
		}
	}
	return place
}