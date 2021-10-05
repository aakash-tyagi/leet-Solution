package main

import (
	"fmt"
)

func selfDividingNumbers(left int, right int) []int {
	res := []int{}

	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			res = append(res, i)
		}
	}

	return res
}

func isSelfDividing(i int) bool {
	r := i

	for ; r != 0; r /= 10 {
		d := r % 10
		if d == 0 || i%d != 0 {
			return false
		}
	}

	return true
}

func main() {

	fmt.Println(selfDividingNumbers(1, 20))
}
