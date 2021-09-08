package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	for n >= 2 {
		if n%2 == 0 {
			n = n / 2
			fmt.Println(n)
			if n == 1 {
				return true
			}
		} else {
			return false
		}
	}

	return false

}

func main() {

	fmt.Println(isPowerOfTwo(64))
}
