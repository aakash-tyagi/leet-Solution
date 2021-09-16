package main

import "fmt"

func lengthOfLastWord(s string) int {

	a := 0
	b := 0
	if s == "" {
		return 0
	}

	l := len(s)

	for i := l - 1; i >= 0; i-- {
		if string(s[i]) == " " {

			a++
		} else {
			break
		}

	}

	for i := l - 1 - a; i >= 0; i-- {
		if string(s[i]) == " " {
			break
		} else {

			b++
		}

	}

	return b
}

func main() {
	a := "how are you"
	fmt.Println(lengthOfLastWord(a))

}
