package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {

	mapS, mapT := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mapS[s[i]]++

	}
	for i := 0; i < len(t); i++ {
		mapT[t[i]]++

	}

	for k, v := range mapT {
		if mapS[k] != v {
			return k

		}
	}
	return t[0]

}

func main() {
	a := findTheDifference("hehehaaa", "hehehaaaw")
	fmt.Println(a)
}
