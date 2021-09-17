package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {

	res := ""
	carry := byte(0)
	i := len(num1) - 1
	j := len(num2) - 1

	for i >= 0 || j >= 0 {
		sum := byte(0)

		if i >= 0 {
			sum += num1[i] - '0'
		}
		if j >= 0 {
			sum += num2[j] - '0'
		}

		sum += carry
		carry = sum / 10
		n := sum % 10
		res = string(n+'0') + res
		i--
		j--
	}
	if carry > 0 {
		res = string(carry+'0') + res
	}

	return res
}

func main() {
	fmt.Println(addStrings("123", "34"))
}
