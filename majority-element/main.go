package main

import (
	"fmt"
)

func majorityElement(nums []int) int {

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++

		if m[v] > len(nums)/2 {
			return v
		}
	}

	return 0
}

func main() {
	a := []int{1, 1, 1, 2}
	fmt.Println(majorityElement(a))
}
