package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	lessFunc := func(i, j int) bool { return intervals[i][0] < intervals[j][0] }
	sort.Slice(intervals, lessFunc)
	result := make([][]int, 0, len(intervals))
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if result[len(result)-1][1] >= intervals[i][0] {
			if result[len(result)-1][1] < intervals[i][1] {
				result[len(result)-1][1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}
