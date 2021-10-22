package leetcode

import (
	"sort"
	"strconv"
)

func LargestNumber(nums []int) string {
	greaterFunc := func(i, j int) bool {
        x, y := nums[i], nums[j]
        sx, sy := 10, 10
        for sx <= x {
            sx *= 10
        }
        for sy <= y {
            sy *= 10
        }
        return sy*x+y > sx*y+x
	}
	sort.Slice(nums, greaterFunc)
	result := ""
	for _, v := range nums {
		result += strconv.Itoa(v)
	}
	if result[0] == '0' {
		return "0"
	} else {
		return result
	}

}
