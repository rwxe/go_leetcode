package leetcode

import "math"

func maxProfit122_1(prices []int) int {
	//状态机空间优化版本
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp_i_0, dp_i_1 := 0, math.MinInt64
	for i := 0; i < len(prices); i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, dp_i_0-prices[i])
	}
	return dp_i_0
}
func maxProfit122_0(prices []int) int {
	//状态机
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := range dp {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		}
	}
	return dp[len(dp)-1][0]
}
