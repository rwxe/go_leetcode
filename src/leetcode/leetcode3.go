package leetcode

//哈希滑动窗口
func lengthOfLongestSubstring(s string) int {
	exists := make(map[byte]bool, len(s))

	maxLen := 0
	for l, r := 0, 0; l < len(s) && r < len(s); r++ {
		for exists[s[r]] {
			exists[s[l]] = false
			l++
		}
		exists[s[r]] = true
		if currLen := r - l + 1; currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}
