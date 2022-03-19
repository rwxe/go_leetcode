package leetcode

import "math"

func MaxProfit121_2(prices []int) int {
	//状态机DP，空间优化版
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp_i_0, dp_i_1 := 0, math.MinInt64
	//dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
	//dp[i][k][1]=max(dp[i-1][k][1],dp[i-1][k-1][0]-prices[i])
	for i := 0; i < len(prices); i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}
func maxProfit121_1(prices []int) int {
	//状态机DP
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
	//dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
	//dp[i][k][1]=max(dp[i-1][k][1],dp[i-1][k-1][0]-prices[i])
	for i := range dp {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], -prices[i])
		}
	}
	return dp[len(dp)-1][0]
}
func maxProfit121_0(prices []int) int {
	//简单DP
	dpPre, dpCurrent := 0, 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
		if v := prices[i] - minPrice; v > dpPre {
			dpCurrent = v
			dpPre = dpCurrent
		} else {
			dpCurrent = dpPre
		}
	}
	return dpCurrent

}
