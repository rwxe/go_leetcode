package leetcode

import "math"

func findClosestNumber(nums []int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		} else {
			return a
		}
	}
	closestDist := math.MaxInt64
	closestNum := math.MaxInt64
	for i := range nums {
		if dist := abs(nums[i]); dist < closestDist {
			closestDist = dist
			closestNum = nums[i]
		} else if dist == closestDist && nums[i] > closestNum {
			closestNum = nums[i]
		}
	}
	return closestNum
}
