package main

import (
	"fmt"
)

func moveZeroes(nums []int) {

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}

}

func main() {
	a := []int{0, 1, 3, 4, 5, 0, 4, 4}
	moveZeroes(a)
	fmt.Println(a)
}
