package main

import "fmt"

func guessNumber(n int) int {

	l, r := 1, n

	for l < r {
		m := l + (r-l)/2
		if guess(m) == 1 {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func guess(num int) int {
	no := 10

	if num < no {
		return -1
	} else if num > no {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(guessNumber(21))
}
