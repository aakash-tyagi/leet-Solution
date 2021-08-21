package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i <= len(nums)-1; i++ {
		for j := 0; j <= len(nums)-1; j++ {
			if i == j {

			} else {
				if nums[i]+nums[j] == target {
					a := []int{i, j}
					return a
				}
			}

		}
	}
	return nil
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(twoSum(a, 7))
}
