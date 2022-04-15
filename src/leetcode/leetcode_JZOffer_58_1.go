package leetcode

import "strings"

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	//words := strings.Fields(s)
	words := strings.FieldsFunc(s, func(c rune) bool {
		if c == rune(' ') {
			return true
		} else {
			return false
		}
	})
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		if i != 0 {
			result += " "
		}
	}
	return result
}
