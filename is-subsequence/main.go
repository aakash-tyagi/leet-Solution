package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {

	if s == "" {
		return true
	}
	res := len(s)
	for i, j := 0, 0; j < len(t); j++ {
		fmt.Println(i, j)
		if s[i] == t[j] {
			res--
			i++
			if res == 0 {
				return true
			}
		}

	}
	return false

}

func main() {

	fmt.Println(isSubsequence("yv", "abcydv"))
}
