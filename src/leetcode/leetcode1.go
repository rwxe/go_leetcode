package leetcode

import "sort"

// 双循环
func twoSum_0(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}

}

//哈希表
func twoSum_1(nums []int, target int) []int {
	//value:index
	hashMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if vi, ok := hashMap[target-v]; ok {
			return []int{vi, i}
		}
		hashMap[v] = i
	}
	return []int{-1, -1}
}

//排序双指针
func twoSum_2(nums []int, target int) []int {
	//value:index
	newNums := make([][2]int, 0, len(nums))
	for i, v := range nums {
		newNums = append(newNums, [2]int{v, i})
	}
	sort.Slice(newNums, func(i, j int) bool {
		return newNums[i][0] < newNums[j][0]
	})
	for i, j := 0, len(newNums)-1; i < j; {
		if sum := newNums[i][0] + newNums[j][0]; sum == target {
			return []int{newNums[i][1], newNums[j][1]}
		} else if sum < target {
			i++
		} else if sum > target {
			j--
		}
	}
	return []int{-1, -1}
}
