package leetcode

import (
	"strconv"
)

func bt93(ipStr string, path []string, start int, result *[]string) {
	if len(path) == 4 && start == len(ipStr) {
		ipAddr := ""
		for i, v := range path {
			ipAddr += v
			if i < len(path)-1 {
				ipAddr += "."

			}
		}
		*result = append(*result, ipAddr)
		return
	} else if len(path) == 4 && start < len(ipStr) {
		return
	}
	for end := start; end < len(ipStr); end++ {
		if end != start && ipStr[start] == '0' {
			return
		}
		if num, _ := strconv.Atoi(string(ipStr[start : end+1])); num > 255 {
			return
		}
		path = append(path, ipStr[start:end+1])
		bt93(ipStr, path, end+1, result)
		path = path[:len(path)-1]
	}
}

func RestoreIpAddresses(s string) []string {
	result := make([]string, 0)
	bt93(s, []string{}, 0, &result)
	return result
}
