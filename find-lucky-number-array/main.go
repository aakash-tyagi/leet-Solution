package main

import (
	"fmt"
)

func findLucky(arr []int) int {

	mapArr := map[int]int{}
	for _, val := range arr {
		mapArr[val]++
	}

	max := 0
	for key, val := range mapArr {
		if key == val {
			if val > max {
				max = val
			}
		}
	}

	if max == 0 {
		return -1
	}
	return max
}

func main() {
	a := []int{2, 2, 2, 3, 3}
	fmt.Println(findLucky(a))
}
