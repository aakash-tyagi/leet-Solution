package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		sum := numbers[i] + numbers[j]
		if (sum) == target {
			break
		} else if (sum) < target {
			i++
		} else {
			j--
		}
	}
	return []int{i + 1, j + 1}
}

func main() {

	a := []int{1, 2, 3, 5}
	fmt.Println(twoSum(a, 3))
}
