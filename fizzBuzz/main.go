package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {

	arr := []string{}

	for i := 1; i <= n; i++ {
		j := ""
		if i%15 == 0 {
			j = "FizzBuzz"
		} else if i%5 == 0 {
			j = "Buzz"
		} else if i%3 == 0 {
			j = "Fizz"
		} else {
			j = strconv.Itoa(i)
		}

		arr = append(arr, j)

	}

	return arr
}

func main() {
	fmt.Println(fizzBuzz(15))
}
