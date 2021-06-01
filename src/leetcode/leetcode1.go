package leetcode

import (
	"sort"
)

func TwoSum(nums []int, target int) []int {
	// index，num
	sortedIndex := make([][]int, len(nums))
	for i, n := range nums {
		sortedIndex[i] = []int{i, n}
	}
	sort.Slice(sortedIndex, func(i, j int) bool { return sortedIndex[i][1] < sortedIndex[j][1] })
	for i, j := 0, len(nums)-1; i != j; {
		sum := sortedIndex[i][1] + sortedIndex[j][1]
		if sum < target {
			i += 1
		} else if sum > target {
			j -= 1
		} else {
			return []int{sortedIndex[i][0], sortedIndex[j][0]}
		}
	}

	return []int{}
}
