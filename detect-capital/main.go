package main

import (
	"fmt"
	"strings"
)

var a []int

func detectCapitalUse(word string) bool {

	if len(word) == 1 {
		return true
	}

	if allCapital(word) || allSmall(word) {
		return true
	}

	ord := strings.TrimPrefix(word, string(word[0]))

	if (word[0] <= 90 && word[0] >= 65) && allSmall(ord) {
		return true
	}

	return false

}

func allSmall(word string) bool {
	for _, v := range word {
		if v <= 122 && v >= 97 {
			a = append(a, 1)
		} else {
			return false
		}
	}

	return len(a) == len(word)

}

func allCapital(word string) bool {
	for _, v := range word {
		if v <= 90 && v >= 65 {
			a = append(a, 1)
		} else {
			return false
		}
	}
	return len(a) == len(word)
}

func main() {
	a := "ggG"
	fmt.Println(detectCapitalUse(a))
}
