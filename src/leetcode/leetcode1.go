package leetcode

import (
	"sort"
)

func TwoSum_0(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}
	return []int{-1, -1}
}

func TwoSum_1(nums []int, target int) []int {
	// target-x:index
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := hashMap[target-nums[i]]; ok {
			return []int{index, i}
		}
		hashMap[nums[i]] = i
	}
	return []int{-1, -1}

}

func TwoSum_2(nums []int, target int) []int {
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
