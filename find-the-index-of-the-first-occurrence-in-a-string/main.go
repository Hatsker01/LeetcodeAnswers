package main

import (
	"fmt"
)

func main() {
	str := "babbbbbabb"
	fmt.Println(strStr(str, "bbab"))
}

func strStr(haystack string, needle string) int {
	var (
		// index int = 0
		j int = 0
	)
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		for k := i; k < len(haystack); k++ {
			if haystack[k] == needle[j] {
				if j == len(needle)-1 {
					return k - len(needle) + 1
				}
				j++
			} else {
				j = 0
				break
				// if haystack[k] == needle[0]{
				// 	j=1
				// }

			}
		}
	}
	return -1
}
