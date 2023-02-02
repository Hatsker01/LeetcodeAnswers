package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {
	n := num
	sum := 0
	for num >= 10 {
		for n > 0 {

			sum += n % 10
			n /= 10
		}
		num = sum
		n = sum
		sum=0
		
	}
	return n
}
