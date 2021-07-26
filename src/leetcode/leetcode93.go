package leetcode

import (
	"strconv"
)

func bt93(ipStr string, path []string, pos int, result *[]string) {
	if len(path) == 4 && pos == len(ipStr) {
		ipAddr := ""
		for i, v := range path {
			ipAddr += v
			if i < len(path)-1 {
				ipAddr += "."
			}
		}
		*result = append(*result, ipAddr)
		return
	} else if len(path) == 4 && pos < len(ipStr) {
		return
	}
	for i := pos; i < len(ipStr); i++ {
		if i != pos && ipStr[pos] == '0' {
			return
		}
		if num, _ := strconv.Atoi(string(ipStr[pos : i+1])); num > 255 {
			return
		}
		path = append(path, ipStr[pos:i+1])
		bt93(ipStr, path, i+1, result)
		path = path[:len(path)-1]
	}
}

func RestoreIpAddresses(s string) []string {
	result := make([]string, 0)
	bt93(s, []string{}, 0, &result)
	return result
}
