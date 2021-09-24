package main

import (
	"fmt"
)

func distributeCandies(candyType []int) int {

	mapCandy := make(map[int]int)

	for _, j := range candyType {
		mapCandy[j]++
	}

	typeCount := 0
	for range mapCandy {

		typeCount++
		if typeCount == len(candyType)/2 {
			break
		} else {
			continue
		}

	}
	return typeCount
}

func main() {
	a := []int{1, 1, 1, 1}
	fmt.Println(distributeCandies(a))
}
