package leetcode

import (
	"math"
)

func SupplyWagon(s []int) []int {
	var sumCount, minL, minR int
	for sumCount != len(s)/2-1 {
		l, r := 0, 1
		minSum := math.MaxInt64
		if sumCount != 0 {

			s[minL] = s[minL] + s[minR]
			s[minR] = 0
		}
		sumCount = 0
		minL, minR = l, r
		for r < len(s) {
			for ; r < len(s) && s[r] == 0; r++ {
			}
			for ; l < len(s) && s[l] == 0; l++ {
			}
			if l < r && r < len(s) {
				if sum := s[l] + s[r]; sum < minSum {
					minSum = sum
					minL, minR = l, r
				}
				//fmt.Println(l, r, minL, minR, sumCount, s)
				sumCount++
				l++
				r++
			}
		}
	}
	//fmt.Println(s)
	result := make([]int, 0, len(s)/2)
	for _, v := range s {
		if v != 0 {

			result = append(result, v)
		}
	}
	return result
}
