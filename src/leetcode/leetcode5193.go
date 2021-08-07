package leetcode

import (
	"strings"
)

func MakeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	arr := []byte{}
	count := 1
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 || s[i] != s[i+1] {
			if count > 2 {
				count = 2
			}
			arr = append(arr, []byte(strings.Repeat(string(s[i]), count))...)
			count = 1
		} else {
			count++
		}
	}
	return string(arr)

}
