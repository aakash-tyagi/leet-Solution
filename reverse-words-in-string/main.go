package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	ss := strings.Fields(s)
	fmt.Println(ss)

	reverse(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func reverse(s *[]string, i int, j int) {

	for i < j {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		i++
		j--
	}
}

func main() {
	a := " wats up Boys"
	fmt.Println(reverseWords(a))
}
