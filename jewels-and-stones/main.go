package main

import (
	"fmt"
)

func numJewelsInStones(jewels string, stones string) int {

	mapStones := make(map[rune]int)

	for _, stone := range stones {
		mapStones[stone]++
	}

	res := 0

	for _, jewel := range jewels {
		if j, found := mapStones[jewel]; found {
			res += j

		}
	}

	return res
}

func main() {
	a := "aA"
	b := "aAAbbbb"
	fmt.Println(numJewelsInStones(a, b))
}
