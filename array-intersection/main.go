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
		m[j]++					// counting number of paticluar int in ARRAY
	}
	
	for _,j := range nums2{
		if m[j] > 0 {
			res = append(res,j) // creating array of intersecting number
			m[j]  = 0
		}
	
	}
	
	return res

}




func main() {
	a := []int{0, 1, 2, 2, 3, 4, 5, 4, 4}
	b := []int{0, 2, 2, 3}
	fmt.Println(intersection(a, b))
}
