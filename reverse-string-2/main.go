package main

import (
	"fmt"
)

func reverseStr(s string, k int) string {

	res := ""
	lenS := len(s)
	if s == "" {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	if k < lenS {
		a := s[:k]
		b := s[k:]
		res = reverse(a)
		res = res + b

		return res
	} else {
		return reverse(s)
	}

}

func reverse(s string) string {

	res := ""
	lenS := len(s)
	arr := make([]string, lenS)
	for i, v := range s {
		arr[i] = string(v)
	}

	for i, j := 0, lenS-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for _, val := range arr {
		res += val
	}

	return res
}

func main() {

	a := "abcdefg"
	fmt.Println(reverseStr(a, 2))
}
