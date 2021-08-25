package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	ar := []byte(a)
	br := []byte(b)
	res := ""

	i := len(ar) - 1
	j := len(br) - 1
	currSum := byte(0)
	add := byte(0)

	for i >= 0 || j >= 0 || add > 0 {
		dar := byte(0)
		dbr := byte(0)

		if i >= 0 {
			dar = ar[i] - '0'
		}
		if j >= 0 {
			dbr = br[j] - '0'
		}

		currSum = dar + dbr + add
		add = currSum >> 1

		res = string(currSum%2+'0') + res

		i--
		j--
	}

	return res
}

func main() {

	a := "111"
	b := "11"
	fmt.Println(addBinary(a, b))
}
