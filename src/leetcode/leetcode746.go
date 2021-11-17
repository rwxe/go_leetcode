package leetcode

func MinCostClimbingStairs(cost []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	dp := make([]int, len(cost)+1)
	dp[0] = 0 // 实际上在Go中会自动初始化为0
	dp[1] = 0
	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[len(dp)-1]

}
func MinCostClimbingStairs_1(cost []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	dp0 := 0
	dp1 := 0
	dpCurrent := 0
	for i := 2; i < len(cost)+1; i++ {
		//dp[i] 取决于 dp[i-2],dp[i-1],cost[i-2],cost[i-1]
		dpCurrent = min(dp0+cost[i-2], dp1+cost[i-1])
		dp0, dp1 = dp1, dpCurrent
	}
	return dpCurrent

}
