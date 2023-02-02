package main

import "fmt"

func main(){
	fmt.Println(isUgly(14))
}

func isUgly(n int) bool {
    for n>=1{
		for n%2==0{
			n/=2
		}
		for n%3==0{
			n/=3
		}
		for n%5==0{
			n/=5
		}
		break
	}
	if n==1{
		return true
	}
	return false
}