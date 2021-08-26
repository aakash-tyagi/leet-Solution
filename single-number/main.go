package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func main() {
	a := []int{1, 2, 2, 3, 3, 1, 5}
	fmt.Println(singleNumber(a))
}
