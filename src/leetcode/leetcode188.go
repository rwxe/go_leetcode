package leetcode

import "math"

func maxProfit188_0(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	if k > n/2 {
		k = n / 2
	}
	dp := make([][][]int, len(prices))
	maxFunc := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	//dp[-1][*][0]=0
	//dp[-1][*][1]=-Inf
	//dp[*][0][0]=0
	//dp[*][0][1]=-Inf
	for i := range dp {
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt64
	}
	for i := 0; i < len(prices); i++ {
		for j := k; j > 0; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
			} else {
				dp[i][j][0] = maxFunc(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[i][j][1] = maxFunc(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])

			}
		}
	}
	return dp[len(dp)-1][k][0]

}
