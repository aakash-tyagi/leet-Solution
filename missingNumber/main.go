package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}

	return len(nums)

}

func main() {
	a := []int{0, 1, 2, 4}
	fmt.Println(missingNumber(a))
}
