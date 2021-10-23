package main

import (
	"fmt"
)

func smallerNumbersThanCurrent(nums []int) []int {

	res := make([]int, len(nums))
	for i, val := range nums {
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] < val {
				count++
			}
		}
		res[i] = count
	}
	return res
}

func main() {
	a := []int{1, 2, 2, 3, 4}
	fmt.Println(smallerNumbersThanCurrent(a))
}
