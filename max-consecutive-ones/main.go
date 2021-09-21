package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {

	res := 0
	con := 0
	for _, val := range nums {

		if val == 1 {
			con++
		}
		if con > res {
			res = con

		}
		if val == 0 {
			con = 0
		}
	}
	return res
}

func main() {
	a := []int{1, 1, 1, 1, 1, 0, 1, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(a))
}
