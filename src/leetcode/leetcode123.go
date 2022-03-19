package leetcode

import (
	"math"
)

func MaxProfit123_2(prices []int) int {
	//这个和版本1代码是一样的
	//但是变量名不同
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	fstBuy, secBuy := math.MinInt64, math.MinInt64
	fstSell, secSell := 0, 0
	for i := 0; i < len(prices); i++ {

		fstBuy = max(fstBuy, -prices[i])
		fstSell = max(fstSell, fstBuy+prices[i])

		secBuy = max(secBuy, fstSell-prices[i])
		secSell = max(secSell, secBuy+prices[i])
	}
	return secSell

}
func MaxProfit123_1(prices []int) int {
	//空间优化版本
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	//dp[i][1][0],dp[i][1][0],dp[i][2][0],dp[i][2][0]
	dp_i_1_0, dp_i_1_1 := 0, math.MinInt64
	dp_i_2_0, dp_i_2_1 := 0, math.MinInt64
	for i := 0; i < len(prices); i++ {

		dp_i_1_1 = max(dp_i_1_1, -prices[i])
		dp_i_1_0 = max(dp_i_1_0, dp_i_1_1+prices[i])

		dp_i_2_1 = max(dp_i_2_1, dp_i_1_0-prices[i])
		dp_i_2_0 = max(dp_i_2_0, dp_i_2_1+prices[i])
	}
	return dp_i_2_0

}
func maxProfit123_0(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	K := 2
	dp := make([][][]int, len(prices))
	for i := range dp {
		//考虑到K可以为0,所以K必须加一
		dp[i] = make([][]int, K+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	//dp[-1][*][0]=dp[*][0][0]=0
	//dp[-1][*][1]=dp[*][0][1]=-Inf
	//
	//dp[i][k][1]=max(dp[i-1][k][1],dp[i-1][k-1][0]-prices[i])
	//dp[i][k][0]=max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
	for i := range dp {
		for k := 1; k <= K; k++ {
			if i == 0 {
				dp[i][k][1] = -prices[i]
				dp[i][k][0] = 0
			} else {
				dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			}
		}
	}
	return dp[len(dp)-1][K][0]
}
