package leetcode

func IsPalindrome125(s string) bool {
	// printable 48-122
	for i, j := 0, len(s)-1; i < j; {
		if s[i] <= 47 || s[i] >= 123 || 58 <= s[i] && s[i] <= 64 || 91 <= s[i] && s[i] <= 96 {
			i++
			continue
		}
		if s[j] <= 47 || s[j] >= 123 || 58 <= s[j] && s[j] <= 64 || 91 <= s[j] && s[j] <= 96 {
			j--
			continue
		}
		if 47 <= s[i] && s[i] <= 57 {
			if s[i] != s[j] {
				return false
			}
		} else {
			if s[i]-s[j] != 32 && s[i]-s[j] != 224 && s[i]-s[j] != 0 {
				return false
			}
		}
		i++
		j--

	}
	return true

}
