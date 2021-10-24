package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	res := [][]int{}
	count := map[int]int{}

	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	var uniqueNum []int
	for key := range count {
		uniqueNum = append(uniqueNum, key)
	}

	sort.Ints(uniqueNum)

	for i := 0; i < len(uniqueNum); i++ {
		if uniqueNum[i]*4 == target && count[uniqueNum[i]] > 3 {
			res = append(res, []int{uniqueNum[i], uniqueNum[i], uniqueNum[i], uniqueNum[i]})
		}
		for j := i + 1; j < len(uniqueNum); j++ {
			if (uniqueNum[i]*3+uniqueNum[j] == target) && count[uniqueNum[i]] > 2 {
				res = append(res, []int{uniqueNum[i], uniqueNum[i], uniqueNum[i], uniqueNum[j]})
			}
			if (uniqueNum[i]*2+uniqueNum[j]*2 == target) && (count[uniqueNum[i]] > 1 && count[uniqueNum[j]] > 1) {
				res = append(res, []int{uniqueNum[i], uniqueNum[i], uniqueNum[j], uniqueNum[j]})
			}
			if (uniqueNum[j]*3+uniqueNum[i] == target) && count[uniqueNum[j]] > 2 {
				res = append(res, []int{uniqueNum[i], uniqueNum[j], uniqueNum[j], uniqueNum[j]})
			}
			for k := j + 1; k < len(uniqueNum); k++ {
				if (uniqueNum[i]*2+uniqueNum[j]+uniqueNum[k] == target) && count[uniqueNum[i]] > 1 {
					res = append(res, []int{uniqueNum[i], uniqueNum[i], uniqueNum[j], uniqueNum[k]})
				}
				if (uniqueNum[j]*2+uniqueNum[i]+uniqueNum[k] == target) && count[uniqueNum[j]] > 1 {
					res = append(res, []int{uniqueNum[i], uniqueNum[j], uniqueNum[j], uniqueNum[k]})
				}
				if (uniqueNum[k]*2+uniqueNum[j]+uniqueNum[i] == target) && count[uniqueNum[k]] > 1 {
					res = append(res, []int{uniqueNum[i], uniqueNum[j], uniqueNum[k], uniqueNum[k]})
				}
				c := target - uniqueNum[i] - uniqueNum[j] - uniqueNum[k]
				if c > uniqueNum[k] && count[c] > 0 {
					res = append(res, []int{uniqueNum[i], uniqueNum[j], uniqueNum[k], c})
				}

			}
		}

	}

	return res
}

func main() {
	a := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(a, 0))
}
