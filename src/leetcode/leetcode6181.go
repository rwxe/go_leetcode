package leetcode

func LongestContinuousSubstring(s string) int {
	maxLen := 0
	currLen := 0
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i]-s[i-1] == 1 {
			currLen++
		} else {
			currLen = 1
		}
		maxLen = max(currLen, maxLen)
	}
	return maxLen
}
