package main

import "fmt"

func rob(nums []int) int {

	l := len(nums)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	dp := make([]int, l)

	dp[0], dp[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[l-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func main() {
	a := []int{1, 2, 3, 1}
	fmt.Println(rob(a))
}
