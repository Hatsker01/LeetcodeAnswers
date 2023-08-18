package main

func main() {

}

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var count int
	for _, val := range hours {
		if val <= target {
			count++
		}
	}
	return count
}
