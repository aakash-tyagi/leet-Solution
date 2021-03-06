package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	res, record := []int{}, map[int]int{}
	for i, v := range nums2 {
		record[v] = i
	}
	fmt.Println(record)
	for i := 0; i < len(nums1); i++ {
		flag := false
		for j := record[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				flag = true
				break
			}
		}
		if flag == false {
			res = append(res, -1)
		}
	}
	return res
}

func main() {
	a := []int{4, 1, 2}
	b := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(a, b))
}
