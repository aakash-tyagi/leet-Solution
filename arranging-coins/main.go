package main

import (
	"fmt"
)

func arrangeCoins(n int) int {

	counter := 0
	i := 1
	for n > 0 {
		n = n - i
		if n < 0 {
			return counter
		}
		counter++
		i++
	}
	return counter
}

func main() {
	fmt.Println(arrangeCoins(1))
}
