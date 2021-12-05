package main

import "fmt"

func findDuplicates(nums []int) []int {

	mapD := make(map[int]int)
	arrD := []int{}

	for _, val := range nums {
		mapD[val]++
	}

	for key, val := range mapD {
		if val == 2 {
			arrD = append(arrD, key)
		}
	}

	return arrD
}

func main() {
	a := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(a))
}
