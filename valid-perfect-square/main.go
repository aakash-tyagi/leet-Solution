package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {

	for i := 2; i <= num/2; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(1022))
}
