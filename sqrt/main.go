package main

import (
	"fmt"
)

func mySqrt(x int) int {

	for i := 0; i <= x; i++ {
		if i*i == x {
			return i
		}
		if i*i > x {
			return i - 1
		}
	}
	return 0

}

func main() {
	fmt.Println(mySqrt(8))
}
