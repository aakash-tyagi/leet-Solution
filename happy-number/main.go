package main

import (
	"fmt"
)

func isHappy(n int) bool {

	res := 0
	num := n
	mapN := make(map[int]int)

	for {
		for num != 0 {
			res += (num % 10) * (num % 10)
			num = num / 10
		}

		if res == 1 {
			return true
		}

		_, ok := mapN[res]

		if !ok {
			mapN[res] = res
			num = res
			res = 0
			continue

		} else {
			return false
		}
	}

}

func main() {

	fmt.Println(isHappy(19))
}
