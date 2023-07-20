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
	numsIndex := make([][2]int, 0, len(nums))
	for i, v := range nums {
		numsIndex = append(numsIndex, [2]int{v, i})
	}
	sort.Slice(numsIndex, func(i, j int) bool {
		return numsIndex[i][0] < numsIndex[j][0]
	})
	for i, j := 0, len(numsIndex)-1; i < j; {
		if sum := numsIndex[i][0] + numsIndex[j][0]; sum == target {
			return []int{numsIndex[i][1], numsIndex[j][1]}
		} else if sum < target {
			i++
		} else if sum > target {
			j--
		}
	}
	return []int{-1, -1}
}
