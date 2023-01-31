package main

import "fmt"

func main() {

	fmt.Println(getRow(4))
}
func getRow(rowIndex int) []int {
	var num []int
	rowIndex++
	for i := 0; i < rowIndex; i++ {
		if i == 0 || i == rowIndex-1 {
			num = append(num, 1)
		} else {
			num = append(num, (num[i-1]*(rowIndex-i))/i)
		}
	}
	return num
}
