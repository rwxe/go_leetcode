package leetcode

func lengthOfLongestSubstring_1(s string) int {
	//window left
	l := 0
	showedUp := make(map[byte]bool)
	maxLen := 0
	for r := 0; r < len(s); r++ {
		for showedUp[s[r]] {
			delete(showedUp, s[l])
			l++
		}
		showedUp[s[r]] = true
		if maxLen < r-l+1 {
			maxLen = r - l + 1
		}
	}
	return maxLen
}
func lengthOfLongestSubstring_0(s string) int {
	window := make([]byte, 0, len(s))
	showedUp := make(map[byte]bool)
	maxLen := 0
	for i := 0; i < len(s); i++ {
		for showedUp[s[i]] {
			delete(showedUp, window[0])
			window = window[1:]
		}
		window = append(window, s[i])
		showedUp[s[i]] = true
		if maxLen < len(window) {
			maxLen = len(window)
		}
	}
	return maxLen
}
