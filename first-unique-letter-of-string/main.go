package main

import "fmt"

func firstUniqChar(s string) int {
	a := make([]int, 26)

	for i := 0; i < len(s); i++ {
		a[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if a[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {

	fmt.Println(firstUniqChar("wwattf"))

}
