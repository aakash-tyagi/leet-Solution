package main

import (
	"fmt"
)

func search(nums []int, target int) int {

	for i, val := range nums {
		if target == val {
			return i
		}
	}
	return -1
}

func main() {
	a := []int{1, 3, 4, 6}
	fmt.Println(search(a, 6))
}
