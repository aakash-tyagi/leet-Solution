package main

import (
	"fmt"
)

func generate(numRows int) [][]int {

	a := [][]int{}

	for i := 0; i < numRows; i++ {
		a = append(a, []int{})
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				a[i] = append(a[i], 1)

			} else {
				a[i] = append(a[i], a[i-1][j-1]+a[i-1][j])

			}

		}
	}

	return a
}

func main() {

	fmt.Println(generate(10))
}
