package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
}

func isPowerOfTwo(n int) bool {
	var number int = 0
	for i := 0; number < n; i++ {
		if int(math.Pow(2, float64(i))) == n {
			return true
		}
		number=int(math.Pow(2,float64(i)))
	}
	return false
}
