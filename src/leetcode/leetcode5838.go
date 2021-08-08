package leetcode

func IsPrefixString(s string, words []string) bool {
	i := 0
	for _, word := range words {
		seg := i + len(word)
		if i < len(s) {
			for j := 0; i < seg; i, j = i+1, j+1 {
				if i >= len(s) {
					return false
				}
				if s[i] != word[j] {
					return false
				}
			}
		}
	}
	if i == len(s) {
		return true

	} else {
		return false
	}

}
