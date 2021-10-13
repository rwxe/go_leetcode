package leetcode

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			j, k := i, 0
			for ; k < len(needle) && haystack[j] == needle[k]; j, k = j+1, k+1 {}
			if j==len(needle)+i{
				return i
			}
		}
	}
	return -1
}
