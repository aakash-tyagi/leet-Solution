package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {

	sort.Ints(nums)
	res := 0

	for i := 0; i < len(nums); i = i + 2 {
		res += nums[i]
	}

	return res
}

func main() {
	a := []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(a))
}
