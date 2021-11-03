package main

import (
	"fmt"
)

func countKDifference(nums []int, k int) int {

	counter := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]-nums[j] == k || nums[i]-nums[j] == -k {
				counter++
			}
		}
	}
	return counter / 2
}

func main() {
	a := []int{3, 2, 1, 5, 4}
	fmt.Println(countKDifference(a, 2))
}
