package leetcode

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		// if s[l]==s[r]{continue}

		if s[l] != s[r] {
			//delte Left
			//delte Right
			dL, dR := true, true
			//try delete Left
			for i, j := l+1, r; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					dL = false
					break
				}
			}
			//try delete Right
			for i, j := l, r-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					dR = false
					break
				}
			}
			if dL || dR {
				return true
			} else {
				return false
			}
		}
	}
	return true
}
