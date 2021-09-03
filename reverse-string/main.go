package main

import (
	"fmt"
)

func reverseString(s []string) {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func main() {
	a := []string{"a", "b", "c", "d"}
	reverseString(a)
	fmt.Println(a)
}
