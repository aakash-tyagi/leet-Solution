package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {

	for i := len(stones) - 1; i >= 1; i-- {
		sort.Ints(stones)
		stones[i-1] = stones[i] - stones[i-1]
		stones = stones[:i]
		fmt.Println(stones)
	}

	return stones[0]
}

func main() {
	a := []int{1}
	fmt.Println(lastStoneWeight(a))
}
