package leetcode

import (
	"sort"
	"unicode"
)

func Expand(s string) []string {
	candidates := make([][]byte, 0)
	temp := make([]byte, 0)
	inBraces := false
	result := make([]string, 0)
	for _, c := range s {
		if c == '{' {
			inBraces = true
		} else if c == '}' {
			inBraces = false
			if len(temp) != 0 {
				sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
				candidates = append(candidates, temp)
				temp = make([]byte, 0)
			}
		} else if inBraces {
			if unicode.IsLetter(c) {
				temp = append(temp, byte(c))
			}
		} else {
			candidates = append(candidates, []byte{byte(c)})
		}
	}
	var dfs func(int, []byte)
	dfs = func(level int, path []byte) {
		if len(path) == len(candidates) {
			result = append(result, string(path))
			return
		}
		for _, c := range candidates[level] {
			path = append(path, byte(c))
			dfs(level+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []byte{})
	return result
}
