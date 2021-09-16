package main

import (
	"fmt"
	"strconv"
)

func hammingWeight(num uint32) int {

	s := strconv.FormatUint(uint64(num), 2)

	sum := 0
	for _, j := range s {

		if string(j) == "1" {
			sum += 1 //counting number of 1`s in string
		}
	}

	return sum
}

func main() {

	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
