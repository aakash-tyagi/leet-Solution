package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {

	mapNums := make(map[int]int)
	for _, val := range nums {
		mapNums[val]++
	}

	keyArr := []int{}
	for k := range mapNums {
		keyArr = append(keyArr, k)
	}

	sort.Ints(keyArr)

	res := 0
	if len(keyArr) >= 3 {
		res = keyArr[len(keyArr)-3]
	} else if len(keyArr) == 2 {
		res = keyArr[len(keyArr)-1]
	} else {
		res = keyArr[0]
	}
	return res
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(thirdMax(a))
}
