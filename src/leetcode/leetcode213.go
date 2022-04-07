package leetcode

func rob2_0(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp := func(start, end int) int {
		dp_i, dp_i_1, dp_i_2 := 0, 0, 0
		for i := end; i >= start; i-- {
			dp_i = max(
				nums[i]+dp_i_2,
				dp_i,
			)
			dp_i_1, dp_i_2 = dp_i, dp_i_1
		}
		return dp_i
	}
	return max(
		dp(0, len(nums)-2),
		dp(1, len(nums)-1),
	)
}
