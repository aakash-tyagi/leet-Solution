package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {

	if n == 0 {
		return true
	}

	for i, c := 0, 0; i < len(flowerbed); i++ {
		var prev, next int
		if i != 0 {
			prev = flowerbed[i-1]
		}
		if i != len(flowerbed)-1 {
			next = flowerbed[i+1]
		}

		if (prev == 0 && next == 0) && (flowerbed[i] == 0) {
			flowerbed[i] = 1
			c++
		}

		if c == n {
			return true
		}
	}

	return false
}

func main() {
	a := []int{1, 0, 0, 0, 1}
	fmt.Println(canPlaceFlowers(a, 1))
}
