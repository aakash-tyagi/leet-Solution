package main

import (
	"fmt"
)

func climbStairs(n int) int {
	nums := make([]int, n+1)
	nums[0] = 1
	nums[1] = 1

	for i := 2; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2]

	}

	return nums[n]
}

func main() {

	fmt.Println(climbStairs(4))
}
