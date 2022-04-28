package leetcode

func lengthOfLIS_GREEDY_0(nums []int) int {
	tower := make([]int, len(nums)+1)
	currLen := 1
	tower[currLen] = nums[0]
	for _, v := range nums {
		if tower[currLen] < v {
			currLen++
			tower[currLen] = v
		} else {
			l, r := 1, currLen
			replacePos := 0
			for l <= r {
				mid := l + (r-l)/2
				if tower[mid] < v {
					replacePos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			tower[replacePos+1] = v
		}
	}
	return currLen
}
func lengthOfLIS_DP_0(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp := make([]int, len(nums))
	maxLen := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen

}
