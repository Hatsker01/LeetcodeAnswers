package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	areaMax := 0
	left := 0
	right := len(height) - 1

	for right != left {
		newArra := math.Min(float64(height[left]), float64(height[right])) * (float64(right) - float64(left))
		if math.Min(float64(height[left]), float64(height[right]))*(float64(right)-float64(left)) > float64(areaMax) {
			areaMax = int(newArra)
		}
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}
	// }
	return areaMax
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
