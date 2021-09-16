package main

import "fmt"

func isPowerOfThree(s int) bool {
	for s >= 3 {
		if s%3 == 0 {
			s = s / 3
			if s == 1 {
				return true
			}
		} else {
			return false
		}
	}

	return false
}

func main() {
	fmt.Println(isPowerOfThree(9))
}
