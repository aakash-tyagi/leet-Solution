package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {

	alphaMap := make(map[string]int)

	for i := 0; i < len(magazine); i++ {
		alphaMap[string(magazine[i])]++

	}

	for j := 0; j < len(ransomNote); j++ {
		_, ok := alphaMap[string(ransomNote[j])]
		if !ok {
			return false
		}

		alphaMap[string(ransomNote[j])]--

		if alphaMap[string(ransomNote[j])] < 0 {
			return false
		}

	}

	return true

}

func main() {
	fmt.Println(canConstruct("abcb", "abcd"))
}
