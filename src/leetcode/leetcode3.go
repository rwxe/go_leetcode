package leetcode

//哈希滑动窗口
func lengthOfLongestSubstring(s string) int {
	var maxLen int
	exist := make(map[byte]bool, len(s))
	for l, r := 0, 0; r < len(s); r++ {
		//不需要if，不符合条件不会进入循环
		for ; exist[s[r]]; l++ {
			exist[s[l]] = false
		}
		exist[s[r]] = true
		if currLen := r - l + 1; currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}
