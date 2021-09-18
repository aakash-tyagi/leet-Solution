package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {

	tmp := make(map[int]int)
	temp := []int{}

	for _, v := range nums {
		tmp[v]++
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := tmp[i]; !ok {
			temp = append(temp, i)
		}
	}

	return temp

}

func main() {
	a := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(a))
}
