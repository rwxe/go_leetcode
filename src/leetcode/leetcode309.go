package leetcode

func maxProfit309_0(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	coolDown := 1
	dp := make([][2]int, len(prices))
	//dp[-1][*][0]=dp[*][0][0]=0
	//dp[-1][*][1]=dp[*][0][1]=-Inf
	//
	//dp[i][k][1]=max(dp[i-1][k][1],dp[i-cd][k-1][0]-prices[i])
	//dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
	for i := range dp {
		if i == 0 {
			dp[i][1] = -prices[i]
			dp[i][0] = 0
		} else if i-coolDown == 0 {
			dp[i][1] = max(dp[i-1][1], -prices[i])
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-(1+coolDown)][0]-prices[i])
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		}
	}

	return dp[len(dp)-1][0]
}
