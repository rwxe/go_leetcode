package leetcode

import (
	"math"
)

func MaxProfit188_2(k int, prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	if n := len(prices); n == 0 {
		return 0
	} else if k >= n/2 {
		//k若是大于等于n/2，那可以视为k为正无穷
		return maxProfit122_1(prices)
	}
	//dp[-1][*][0]=dp[*][0][0]=0
	//dp[-1][*][1]=dp[*][0][1]=-Inf
	//
	//dp[i][k][1]=max(dp[i-1][k][1],dp[i-1][k-1][0]-prices[i])
	//dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
	dp_i_k := make([][2]int, k+1)
	for j := range dp_i_k {
		dp_i_k[j][1] = math.MinInt64
		dp_i_k[j][0] = 0 //这步Go其实已经默认做了
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp_i_k[j][1] = max(dp_i_k[j][1], dp_i_k[j-1][0]-prices[i])
			dp_i_k[j][0] = max(dp_i_k[j][0], dp_i_k[j][1]+prices[i])
		}
	}
	return dp_i_k[k][0]

}
func maxProfit188_1(k int, prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	if n := len(prices); n == 0 {
		return 0
	} else if k >= n/2 {
		//k若是大于n/2，那可以视为k为正无穷
		return maxProfit122_1(prices)
	}

	dp := make([][][]int, len(prices))
	for i := range dp {
		//考虑到k为0有特殊含义，因此必须创建k+1个列表
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	//dp[-1][*][0]=dp[*][0][0]=0
	//dp[-1][*][1]=dp[*][0][1]=-Inf
	//
	//dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
	//dp[i][k][1]=max(dp[i-1][k][1],dp[i-1][k-1][0]-prices[i])

	for i := 0; i < len(dp); i++ {
		for j := 1; j <= k; j++ {
			//这里k从小到大或从大到小遍历都可以
			if i == 0 {
				dp[i][j][1] = -prices[i]
				dp[i][j][0] = 0
			} else {
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			}
		}
	}
	return dp[len(dp)-1][k][0]

}
func maxProfit188_0(k int, prices []int) int {
	//未优化
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	if len(prices) == 0 {
		return 0
	}
	dp := make([][][]int, len(prices))
	for i := range dp {
		//考虑到k为0有特殊含义，因此必须创建k+1个列表
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	//dp[-1][*][0]=dp[*][0][0]=0
	//dp[-1][*][1]=dp[*][0][1]=-Inf
	//
	//dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
	//dp[i][k][1]=max(dp[i-1][k][1],dp[i-1][k-1][0]-prices[i])
	for i := 0; i < len(dp); i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[i][j][1] = -prices[i]
				dp[i][j][0] = 0
			} else {
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			}
		}
	}
	return dp[len(dp)-1][k][0]

}
