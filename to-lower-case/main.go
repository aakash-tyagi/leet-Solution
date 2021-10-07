package main

import (
	"fmt"
)

func toLowerCase(s string) string {

	a := []rune(s)

	for i := 0; i < len(a); i++ {
		if a[i] >= 65 && a[i] <= 90 {
			a[i] = a[i] + 32
		}
	}

	return string(a)
}

func main() {
	a := "AZaz"
	fmt.Println(toLowerCase(a))
}
