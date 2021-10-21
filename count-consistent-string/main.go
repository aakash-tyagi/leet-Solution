package main

import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {

	mapAllowed := make(map[rune]int)
	res := 0

	for _, val := range allowed {
		mapAllowed[val]++
	}

	for _, val := range words {
		temp := 0
		for _, sVal := range val {
			if _, ok := mapAllowed[sVal]; !ok {
				break
			} else {
				temp++
			}
		}
		if temp == len(val) {
			res++
		}
	}

	return res

}

func main() {
	a := []string{"cf", "fa", "a", "aaaab"}
	fmt.Println(countConsistentStrings("ab", a))
}
