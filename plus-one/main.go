package main

import (
	"fmt"
)

func plusOne(digits []int) []int {

	nums := make([]int, len(digits))
	copy(nums, digits)
	carry := 1

	for i := len(nums) - 1; i >= 0; i-- {
		num := (nums[i] + carry) % 10

		carry = (nums[i] + carry) / 10

		nums[i] = num

		if carry == 0 {
			return nums
		}

	}

	if carry != 0 {
		nums = append([]int{carry}, nums...)
	}

	return nums

}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}
