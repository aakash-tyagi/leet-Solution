package main

import (
	"fmt"
)

func jump(nums []int) int {

	curJump, jump, furhJump := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {

		if i+nums[i] > furhJump {
			furhJump = i + nums[i]

		}

		if curJump == i {
			jump, curJump = jump+1, furhJump

			if curJump >= len(nums)-1 {
				return jump
			}
		}

	}
	return 0
}

func main() {
	a := []int{1, 1, 3, 1, 4, 6, 3, 2, 1, 1}
	fmt.Println(jump(a))
}
