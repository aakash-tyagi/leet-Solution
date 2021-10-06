package main

import (
	"fmt"
)

func numOfStrings(patterns []string, word string) int {
	res, n, temp := 0, len(word), make(map[string]bool)

	for i := 1; i <= n; i++ {
		for j := 0; j+i <= n; j++ {
			temp[word[j:j+i]] = true
		}
	}

	for _, pattern := range patterns {
		if _, ok := temp[pattern]; ok {
			res++
		}
	}
	return res

}

func main() {
	a := []string{"a", "abc", "bc", "d"}
	b := numOfStrings(a, "abc")
	fmt.Println(b)
}
