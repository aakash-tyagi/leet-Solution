package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {

	mapArr := make(map[int]int)

	for _, j := range nums {
		mapArr[j]++
		val := mapArr[j]
		if val == 2 {
			return true
		}
	}

	return false

}

func main() {
	a := []int{1, 2, 3, 4, 5, 1}
	fmt.Println(containsDuplicate(a))
}
