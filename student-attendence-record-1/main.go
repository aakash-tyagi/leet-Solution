package main

import (
	"fmt"
)

func checkRecord(s string) bool {

	absent := 0
	for _, v := range s {
		if v == 65 {
			absent++
		}
		if absent >= 2 {
			return false
		}
	}

	clate := 0
	for _, v := range s {

		if v == 76 {
			clate++
			if clate >= 3 {
				return false
			}
		} else {
			clate = 0
		}

	}

	return true
}

func main() {

	fmt.Println(checkRecord("ALLLPPPP"))
}
