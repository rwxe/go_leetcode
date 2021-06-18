package leetcode

import "fmt"



func StrStr(haystack string, needle string) int {
	if len(needle)==0{
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {

		if haystack[i] == needle[0] {
			j, k := i, 0
			for ; k < len(needle) && haystack[j] == needle[k]; j, k = j+1, k+1 {

			}
			fmt.Println(i,j,k)
			if i+len(needle) == j {
				return i
			} else {

			}
		}

	}
	return -1
}
