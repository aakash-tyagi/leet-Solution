package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}

	mapNum := map[int]int{}
	for i := 0; i < len(nums); i++ {
		mapNum[nums[i]]++
	}

	var uniqueNum []int

	for key := range mapNum {
		uniqueNum = append(uniqueNum, key)
	}

	sort.Ints(uniqueNum)

	for i := 0; i < len(uniqueNum); i++ {
		if uniqueNum[i]*3 == 0 && mapNum[uniqueNum[i]] > 2 {
			res = append(res, []int{uniqueNum[i], uniqueNum[i], uniqueNum[i]})
		}
		for j := i + 1; j < len(uniqueNum); j++ {
			if (uniqueNum[i]*2+uniqueNum[j] == 0) && mapNum[uniqueNum[i]] > 1 {
				res = append(res, []int{uniqueNum[i], uniqueNum[i], uniqueNum[j]})
			}
			if (uniqueNum[j]*2+uniqueNum[i] == 0) && mapNum[uniqueNum[j]] > 1 {
				res = append(res, []int{uniqueNum[i], uniqueNum[j], uniqueNum[j]})
			}
			c := 0 - uniqueNum[i] - uniqueNum[j]
			if c > uniqueNum[j] && mapNum[c] > 0 {
				res = append(res, []int{uniqueNum[i], uniqueNum[j], c})
			}
		}
	}

	return res
}

func main() {
	a := []int{1, 1, -2}
	fmt.Println(threeSum(a))
}
