package leetcode

// DP
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := range dp {
		dp[i][i] = true
	}
	begin := 0
	maxLen := 1
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			//dp[i][j]= s[i]==s[j] && dp[i+1][j-1]
			//dp[i][j]= s[i]==s[j] && i+1>j-1
			if s[i] == s[j] {
				if i+1 <= j-1 {
					dp[i][j] = dp[i+1][j-1]
				} else {
					dp[i][j] = true
				}
			}
			if currLen := j - i + 1; dp[i][j] && currLen > maxLen {
				maxLen = currLen
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]

}

//暴力
func longestPalindrome_1(s string) string {
	if len(s) < 2 {
		return s
	}

	vaildPalindrome := func(left, right int) bool {
		for i, j := left, right; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	maxLen := 1
	var begin int

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if vaildPalindrome(i, j) {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					begin = i
				}
			}
		}
	}
	return s[begin : begin+maxLen]

}
