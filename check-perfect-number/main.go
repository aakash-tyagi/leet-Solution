package main

import (
	"fmt"
)

func checkPerfectNumber(num int) bool {

	if num == 1 {
		return false
	}
	var sum int
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	if sum == num {
		return true
	} else {
		return false
	}
}

func main() {

	fmt.Println(checkPerfectNumber(8121))
}
