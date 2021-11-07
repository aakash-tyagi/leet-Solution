package main

import "fmt"

func rotate(nums []int, k int) {

	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = result[i]
	}

}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 1)
	fmt.Println(a)
}
