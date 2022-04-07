package leetcode

import "strconv"

func translateNum_1(num int) int {
	numStr := strconv.Itoa(num)
	//dp[i-2],dp[i-1],dp[i]
	dp_i_1, dp_i := 1, 1
	for i := 2; i <= len(numStr); i++ {
		if "10" <= numStr[i-2:i] && numStr[i-2:i] <= "25" {
			dp_i_1, dp_i = dp_i, dp_i+dp_i_1
		} else {
			dp_i_1, dp_i = dp_i, dp_i
		}
	}
	return dp_i
}
func translateNum_0(num int) int {
	numStr := strconv.Itoa(num)
	dp := make([]int, len(numStr)+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= len(numStr); i++ {
		if "10" <= numStr[i-2:i] && numStr[i-2:i] <= "25" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}
