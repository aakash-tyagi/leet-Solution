package main

import (
	"fmt"
)

func longestPalindrome(s string) int {

	lenS := len(s)
	if lenS == 1 {
		return 1
	}
	res := 0
	mapS := make(map[byte]int)

	for i := 0; i < lenS; i++ {
		mapS[s[i]]++
	}

	var counter int = 0

	for _, val := range mapS {
		if val == 1 && counter < 1 {
			res += 1
			counter++
		}

		if val%2 == 0 {

			res += val
		} else if val%2 != 0 && val > 1 {
			res += val

			counter += 1
			if counter > 1 {
				res -= 1
			}

		}

	}

	return res
}

func main() {
	fmt.Println(longestPalindrome("bananas"))
}
