package main

import (
	"fmt"
)

func validPalindrome(s string) bool {
	return helper([]byte(s), 0, len(s)-1, false)
}

func helper(sByte []byte, lo, hi int, hasDeleted bool) bool {

	for lo < hi {
		if sByte[lo] != sByte[hi] {
			if hasDeleted {
				return false
			}
			return helper(sByte, lo+1, hi, true) || helper(sByte, lo, hi-1, true)
		}
		lo++
		hi--
	}
	return true
}

func main() {
	a := "abdda"
	fmt.Println(validPalindrome(a))
}
