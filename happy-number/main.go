package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	var num int = n
	for  num!=2&&num != 1 {
		var a int=0
		for num > 0 {
			a += int(math.Pow(float64((num)%10), 2))
			num /= 10
			fmt.Println(a)
		}
		num=a
	}
	if num == 1 {
		return true
	}
	return false
}
