package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {

	for i, val := range nums {
		nums[i] = val * val
	}

	sort.Ints(nums)

	return nums
}

func main() {
	a := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(a))
}
