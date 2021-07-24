package leetcode

import (
	"strconv"
)

func bt93(ip string, path []string, segStart int, result *[]string) {
	if len(path) == 4 && segStart == len(ip) {
		ipStr := ""
		for i, v := range path {
			ipStr += v
			if i < len(path)-1 {
				ipStr += "."

			}
		}
		*result = append(*result, ipStr)
		return
	} else if len(path) == 4 && segStart < len(ip) {
		return
	}
	for i := segStart; i < len(ip); i++ {
		if i != segStart && ip[segStart] == '0' {
			return
		}
		if num, _ := strconv.Atoi(string(ip[segStart : i+1])); num > 255 {
			return
		}
		path = append(path, ip[segStart:i+1])
		bt93(ip, path, i+1, result)
		path = path[:len(path)-1]
	}
}

func RestoreIpAddresses(s string) []string {
	result := make([]string, 0)
	bt93(s, []string{}, 0, &result)
	return result

}
