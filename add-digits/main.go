package main

import (
	"fmt"
	"strconv"
)

func addDigits(num int) int {

	str := fmt.Sprint(num)
	var res string
	var a int
	var sum int
	if len(str) == 1 {
		return num
	}

	for i, j := range str {
		temp, _ := strconv.Atoi(string(j))
		sum += temp

		res = fmt.Sprint(sum)
		if len(res) == 1 {
			a, _ = strconv.Atoi(res)
		} else if i == len(str)-1 && len(res) > 1 {
			a = addDigits(sum)
		}
	}

	return a

}

func main() {
	fmt.Println(addDigits(3215))
}
