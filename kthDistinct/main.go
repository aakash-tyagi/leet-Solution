package main

import (
	"fmt"
)

func kthDistinct(arr []string, k int) string {

	mapArr := make(map[string]int)
	count := 1

	for _, val := range arr {
		mapArr[val]++
	}

	for _, val := range arr {
		if mapArr[val] == 1 {
			if count == k {
				return val
			} else if count < k {
				count++
			}
		}
	}

	return ""
}

func main() {
	a := []string{"a", "b", "a"}
	k := 3
	fmt.Println(kthDistinct(a, k))
}
