package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {

	if s == "" {
		return 0
	}

	strs := strings.Fields(s)
	return len(strs)
}

func main() {
	a := "hel l o"
	fmt.Println(countSegments(a))
}
