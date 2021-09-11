package main

import (
	"fmt"
	"sort"
)

package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]int)
	for _,j := range nums1{
		m[j]++
	}
	
	for _,j := range nums2{
		if m[j] > 0 {
			res = append(res,j)
			m[j]  = 0
		}
	
	}
	
	return res

}

func main() {
	a := []int{1,1,2,3}
	b := []int{1,1,4}
	fmt.Println(intersection(a,b))
}


func main() {
	a := []int{0, 1, 2, 2, 3, 4, 5, 4, 4}
	b := []int{0, 2, 2, 3}
	fmt.Println(intersection(a, b))
}
