package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(2147483641))
}
func judgeSquareSum(c int) bool {
	for i := 0; i <= int(math.Sqrt(float64(c))); i++ {
		for k := 0; k <=i; k++ {
			if k*k+i*i == c {
				return true
			}
		}
	}
	return false
}
