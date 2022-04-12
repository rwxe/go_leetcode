package leetcode

import "math"

func minimumDeletions(nums []int) int {
	min := func(items ...int) int {
		minItem := items[0]
		for i := range items {
			if items[i] < minItem {
				minItem = items[i]
			}
		}
		return minItem
	}
	theMin, theMax := math.MaxInt64, math.MinInt64
	theMinIndex, theMaxIndex := 0, 0
	for i, v := range nums {
		if v > theMax {
			theMax = v
			theMaxIndex = i
		}
		if v < theMin {
			theMin = v
			theMinIndex = i
		}
	}
	theFirst, theSecond := 0, 0
	if theMinIndex < theMaxIndex {
		theFirst = theMinIndex
		theSecond = theMaxIndex
	} else {
		theFirst = theMaxIndex
		theSecond = theMinIndex

	}
	return min(
		theFirst+1+len(nums)-theSecond,
		theSecond+1,
		len(nums)-theFirst,
	)

}
