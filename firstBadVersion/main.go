package main

import (
	"fmt"
)

func firstBadVersion(n int) int {

	if n == 1 {
		return 1
	}

	start := 0
	end := n
	index := 0

	for {
		index = (start+end)/2 + 1

		if isBadVersion(index) {
			if !isBadVersion(index - 1) {
				return index
			}
			end = index - 1
			continue
		}
		start = index
	}

	return index
}

func main() {
	fmt.Println(firstBadVersion(5))
}

func isBadVersion(n int) bool {
	if n > 4 {
		return true
	}

	return false
}
