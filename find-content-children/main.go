package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	kidpos := 0
	cookiepos := 0

	for kidpos < len(g) && cookiepos < len(s) {
		if s[cookiepos] >= g[kidpos] { // if the size of the cookie is greater than or equal to the childs greed we are moving up the child pointer to the next child
			kidpos++
		}
		cookiepos++
	}

	return kidpos

}

func main() {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	fmt.Println(findContentChildren(g, s))
}
