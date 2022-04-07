package leetcode

func rob_3(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp_i, dp_i_1, dp_i_2 := 0, 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		//dp[i] = max(nums[i]+dp[i+2], dp[i+1])
		dp_i = max(nums[i]+dp_i_2, dp_i_1)
		dp_i_1, dp_i_2 = dp_i, dp_i_1
	}
	return dp_i

}
func rob_2(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}
	return dp[0]

}
func rob_1(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	memory := make([]int, len(nums))
	for i := range memory {
		memory[i] = -1
	}

	var dp func(pos int) int
	dp = func(pos int) int {
		if pos >= len(nums) {
			return 0
		}
		if memory[pos] != -1 {
			return memory[pos]
		}
		result := max(
			dp(pos+1),
			nums[pos]+dp(pos+2),
		)
		memory[pos] = result
		return result
	}
	return dp(0)

}
func rob_0(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	var dp func(pos int) int
	dp = func(pos int) int {
		if pos >= len(nums) {
			return 0
		}
		return max(
			dp(pos+1),
			nums[pos]+dp(pos+2),
		)
	}
	return dp(0)
}
