package leetcode

import "sort"

func longestConsecutive_1(nums []int) int {
	theSet := make(map[int]bool)
	for i := range nums {
		theSet[nums[i]] = true
	}
	maxLen := 0
	for num := range theSet {
		if !theSet[num-1] {
			currNum := num
			currLen := 0
			for ; theSet[currNum]; currNum++ {
				currLen++
			}
			if currLen > maxLen {
				maxLen = currLen
			}
		}
	}
	return maxLen
}
func longestConsecutive_0(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	theSet := make(map[int]struct{})
	for i := range nums {
		theSet[nums[i]] = struct{}{}
	}
	newNums := make([]int, 0, len(theSet))
	for k := range theSet {
		newNums = append(newNums, k)
	}
	sort.Ints(newNums)
	maxLen := 1
	currLen := 1
	for i := 1; i < len(newNums); i++ {
		if newNums[i]-newNums[i-1] == 1 {
			currLen++
		} else {
			currLen = 1
		}
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}
