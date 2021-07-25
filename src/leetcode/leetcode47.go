package leetcode

import "sort"

func bt47(path, nums []int, length int, result *[][]int) {
	if len(path) >= length {
		*result = append(*result, append([]int(nil), path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		newNums := append([]int(nil), nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		bt47(path, newNums, length, result)
		path = path[:len(path)-1]
	}
}
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	result := make([][]int, 0)
	bt47([]int{}, nums, length, &result)
	return result

}
