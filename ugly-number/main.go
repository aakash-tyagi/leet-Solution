package main

import (
	"fmt"
)

func isUgly(n int) bool {
	for i := 2; n > 0 && i < 6; i++ {
		for n%i == 0 {
			n /= i
		}

	}

	return n == 1

}

func main() {
	fmt.Println(isUgly(11))
}
