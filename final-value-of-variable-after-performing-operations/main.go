package main

import "fmt"

func main() {
	fmt.Println(finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"}))
}

func finalValueAfterOperations(operations []string) int {
	var value int = 0
	for _, val := range operations {
		switch val {
		case "X++", "++X":
			value++
		case "--X", "X--":
			value--
		}

	}
	return value
}
