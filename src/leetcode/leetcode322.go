package leetcode

func CoinChangeNR(coins []int, amount int) int {
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
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}

	}
	if dp[len(dp)-1] == amount+1 {
		return -1
	} else {
		return dp[len(dp)-1]
	}
}

func coinChangeR(coins []int, amount int) int {
	memory := make([]int, amount+1)
	for i := range memory {
		memory[i] = -10
	}
	return coinChangeRHelper(memory, coins, amount)
}

func coinChangeRHelper(memory, coins []int, amount int) int {
	if amount == 0 {
		return 0
	} else if amount < 0 {
		return -1
	} else if memory[amount] != -10 {
		return memory[amount]
	}
	result := amount + 1
	for _, coin := range coins {
		subResult := coinChangeRHelper(memory, coins, amount-coin)
		if subResult == -1 {
			continue
		}
		if subResult+1 < result {
			result = subResult + 1
		}
	}
	if result == amount+1 {
		result = -1
	}
	memory[amount] = result
	return memory[amount]
}
