package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countPrimes(10000))
}
func countPrimes(n int) int {
	sum := 0
	count := 0

	for i := 2; i <= n; i++ {
		if i > 100 {
			for k := 1; k <= int(math.Sqrt(float64(i))); k++ {
				if i%k == 0 {
					count++
				}
			}
		} 
		if (i<=100){
			for k := 1; k <= i; k++ {
				if i%k == 0 {
					count++
				}
			}
		}
		if count == 2 && n != i {
			sum++
		}
		count = 0
	}
	return sum
}
