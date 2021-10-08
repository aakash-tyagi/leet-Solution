package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {

	res := 0
	count := 1
	for num != 0 {
		res += (num % 7) * count
		num /= 7
		count *= 10
	}

	return strconv.Itoa(res)
}

func main() {

	fmt.Println(convertToBase7(14))
}
