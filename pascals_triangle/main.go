package main

import "fmt"

func main() {

	fmt.Println(generate(5))
}
func generate(numRows int) [][]int {
	var num [][]int

	for i := 1; i <= numRows; i++ {
		var b []int

		for k := 0; k < i; k++ {

			if k == 0 || k == i-1 {
				b = append(b, 1)
			} else {
				b = append(b, (b[k-1]*(i-k))/k)
			}

		}
		num = append(num, b)
	}
	return num
}
