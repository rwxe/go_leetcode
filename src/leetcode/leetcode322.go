package leetcode

import "math"

// DP
func coinChange_2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := range dp {
		for _, coin := range coins {
			if i >= coin {
				if dp[i-coin] < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] != amount+1 {
		return dp[amount]
	} else {
		return -1

	}

}
func coinChange_1(coins []int, amount int) int {
	memory := make(map[int]int)
	var dp func(int) int
	dp = func(currAmount int) int {
		if _, ok := memory[currAmount]; ok {
			return memory[currAmount]
		} else if currAmount < 0 {
			return -1
		} else if currAmount == 0 {
			return 0
		} else {
			currResult := math.MaxInt64
			for i := range coins {
				if subResult := dp(currAmount - coins[i]); subResult != -1 {
					if subResult < currResult {
						currResult = subResult + 1
					}
				}
			}
			if currResult != math.MaxInt64 {
				memory[currAmount] = currResult
				return currResult
			} else {
				return -1
			}

		}
	}
	return dp(amount)

}

//无记忆递归
func coinChange_0(coins []int, amount int) int {
	var dp func(int) int
	dp = func(currAmount int) int {
		if currAmount < 0 {
			return -1
		} else if currAmount == 0 {
			return 0
		} else {
			currResult := math.MaxInt64
			for i := range coins {
				if subResult := dp(currAmount - coins[i]); subResult != -1 {
					if subResult < currResult {
						currResult = subResult + 1
					}
				}
			}
			if currResult != math.MaxInt64 {
				return currResult
			} else {
				return -1
			}

		}
	}
	return dp(amount)
}
