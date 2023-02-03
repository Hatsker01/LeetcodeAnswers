package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(isPerfectSquare(16))
}

func isPerfectSquare(num int) bool {
    n:=math.Sqrt(float64(num))
	if float64(int(math.Sqrt(float64(num))))==n{
		return true
	}
	return false
}