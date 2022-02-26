package leetcode

func maximumDifference_0(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		nums[i] = nums[i+1] - nums[i]
	}
	maxDelta, delta := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if delta <= 0 && nums[i] > 0 {
			delta = nums[i]
		} else {
			delta += nums[i]
		}
		if maxDelta < delta {
			maxDelta = delta
		}
	}
	return maxDelta

}
func maximumDifference(nums []int) int {
	theMin := nums[0]
	maxDelta := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < theMin {
			theMin = nums[i]
		} else {
			if maxDelta < nums[i]-theMin && nums[i]-theMin > 0 {
				maxDelta = nums[i] - theMin
			}
		}
	}
	return maxDelta
}
