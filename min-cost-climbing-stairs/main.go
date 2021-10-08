package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {

	dp := make([]int, len(cost))

	dp[0], dp[1] = cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	return min(dp[len(dp)-1], dp[len(dp)-2])

}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}

func main() {
	a := []int{1, 10, 11, 1}
	fmt.Println(minCostClimbingStairs(a))
}
