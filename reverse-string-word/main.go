package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	arrWord := strings.Split(s, " ")

	res := ""
	for i, val := range arrWord {
		res += reverse(val)
		if i < len(arrWord)-1 {
			res += " "
		}
	}

	return res
}

func reverse(s string) string {
	arrAlpha := make([]string, len(s))
	for i, val := range s {
		arrAlpha[i] = string(val)
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		arrAlpha[i], arrAlpha[j] = arrAlpha[j], arrAlpha[i]
	}

	res := ""
	for i := 0; i < len(s); i++ {
		res += arrAlpha[i]
	}

	return res
}

func main() {

	a := "yo wats up"

	fmt.Println(reverseWords(a))
}
