package main

import (
	"fmt"
)

func maximumWealth(accounts [][]int) int {

	sum := 0
	for _, val := range accounts {
		temp := 0
		for _, valJ := range val {
			temp += valJ
		}
		if temp > sum {
			sum = temp
		}
	}

	return sum

}

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 4, 3}
	c := [][]int{a, b}
	fmt.Println(maximumWealth(c))
}
