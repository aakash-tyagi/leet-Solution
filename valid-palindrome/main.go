package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {

	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isChar(c byte) bool { // checking for the string if the byte is character
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

func main() {
	a := "abcba"
	fmt.Println(isPalindrome(a))
}
