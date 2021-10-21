package main

import (
	"fmt"
	"sort"
)

func dominantIndex(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	copyNums := make([]int, len(nums))
	_ = copy(copyNums, nums)

	sort.Ints(nums)

	if nums[(len(nums)-1)]/2 >= nums[len(nums)-2] {
		a := nums[(len(nums) - 1)]

		for i := 0; i < len(nums); i++ {
			fmt.Println(i, copyNums[i])
			if copyNums[i] == a {

				return i
			}
		}
	}
	return -1
}

func main() {
	a := []int{1}
	fmt.Println(dominantIndex(a))
}
