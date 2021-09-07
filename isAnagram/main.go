package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	ls, lt, temp := len(s), len(t), make([]int, 26)

	if ls != lt {
		return false
	}

	for i := 0; i < ls; i++ {
		temp[s[i]-'a']++
		temp[t[i]-'a']--

	}

	tmp := 0

	for _, j := range temp {
		if j < 0 {
			return false
		}

		tmp += j
	}

	return tmp == 0
}

func main() {
	fmt.Println(isAnagram("abaaacd", "abaaacd"))
}
