package leetcode

import "math"

func MaxProfit714_1(prices []int, fee int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp_i_0, dp_i_1 := 0, math.MinInt64
	//dp[-1][*][0]=dp[*][0][0]=0
	//dp[-1][*][1]=dp[*][0][1]=-Inf
	//
	//dp[i][k][1]=max(dp[i-1][k][1],dp[i-1][k-1][0]-prices[i]-fee)
	//dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
	for i := range prices {
		if i == 0 {
			dp_i_1 = -prices[i] - fee
			dp_i_0 = 0
		} else {
			dp_i_1 = max(dp_i_1, dp_i_0-prices[i]-fee)
			dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		}
	}

	return dp_i_0
}
func MaxProfit714_0(prices []int, fee int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp := make([][2]int, len(prices))
	//dp[-1][*][0]=dp[*][0][0]=0
	//dp[-1][*][1]=dp[*][0][1]=-Inf
	//
	//dp[i][k][1]=max(dp[i-1][k][1],dp[i-1][k-1][0]-prices[i]-fee)
	//dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
	for i := range dp {
		if i == 0 {
			dp[i][1] = -prices[i] - fee
			dp[i][0] = 0
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		}
	}

	return dp[len(dp)-1][0]
}
