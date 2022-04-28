package leetcode

import "math"

// DP
func coinChange_2(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := range dp {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if result := dp[len(dp)-1]; result != amount+1 {
		return result
	} else {
		return -1
	}
}

// 带记忆的递归
func coinChange_1(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	memory := make(map[int]int)
	var dfs func(amount int) int
	dfs = func(amount int) int {
		if amount == 0 {
			return 0
		} else if amount < 0 {
			return -1
		} else if res, ok := memory[amount]; ok {
			return res
		}
		res := math.MaxInt64
		for _, coin := range coins {
			if subRes := dfs(amount - coin); subRes == -1 {
				continue
			} else {
				res = min(res, subRes+1)
			}
		}
		if res != math.MaxInt64 {
			memory[amount] = res
			return res
		} else {
			return -1
		}
	}
	return dfs(amount)

}

//无记忆递归
func coinChange_0(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	var dfs func(amount int) int
	dfs = func(amount int) int {
		if amount == 0 {
			return 0
		} else if amount < 0 {
			return -1
		}
		res := math.MaxInt64
		for _, coin := range coins {
			if subRes := dfs(amount - coin); subRes == -1 {
				continue
			} else {
				res = min(res, subRes+1)
			}
		}
		if res != math.MaxInt64 {
			return res
		} else {
			return -1
		}
	}
	return dfs(amount)
}
