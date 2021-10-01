package main

import (
	"fmt"
)

func singleNumber(nums []int) int {

	mapNums := make(map[int]int)

	for _, val := range nums {
		mapNums[val]++
	}

	for k, val := range mapNums {
		if val == 1 {
			return k
		}
	}
	return 0

}

func main() {
	a := []int{1, 2, 1, 2, 4}
	fmt.Println(singleNumber(a))
}
