package leetcode

func findRepeatNumber_2(nums []int) int {
	for i := range nums {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1

}

// 用哈希表
func findRepeatNumber_0(nums []int) int {
	numMap := make(map[int]int)
	for i := range nums {
		if _, ok := numMap[nums[i]]; ok {
			return nums[i]
		}
		numMap[nums[i]]++

	}
	return -1
}

// 用数组
func findRepeatNumber_1(nums []int) int {
	numMap := make([]int, len(nums))
	for i := range nums {
		if numMap[nums[i]] != 0 {
			return nums[i]
		}
		numMap[nums[i]]++
	}
	return -1
}
