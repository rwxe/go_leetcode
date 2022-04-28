package leetcode

func change2_2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[len(dp)-1]
}

//DFS 会超时
func change2(amount int, coins []int) int {
	result := 0
	var dfs func(currAmount, pos int)
	dfs = func(currAmount, pos int) {
		if currAmount == 0 {
			result++
		} else if currAmount < 0 {
			return
		}
		for i := pos; i < len(coins); i++ {
			dfs(currAmount-coins[i], i)
		}
	}
	dfs(amount, 0)
	return result
}
