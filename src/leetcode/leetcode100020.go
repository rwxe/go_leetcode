package leetcode

func MinLengthAfterRemovals(nums []int) int {
	i, j := 0, 1
	for i < len(nums)-1 && j < len(nums) {
		if nums[i] < nums[j] &&
			nums[i] != 0 && nums[j] != 0 {
			nums[i], nums[j] = 0, 0
			i++
		}
		for i < len(nums)-1 && nums[i] == 0 {
			i++
		}
		j++
		if i == j {
			j++
		}
	}
	count := 0
	for _, v := range nums {
		if v != 0 {
			count++
		}

	}
	return count

}
