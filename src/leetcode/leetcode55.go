package leetcode

func canJump(nums []int) bool {
	end := len(nums) - 1
	reachable := 0
	for i := 0; i < len(nums); i++ {
		if i <= reachable {
			currReachable := nums[i] + i
			if currReachable >= end {
				return true
			}
			if currReachable > reachable {
				reachable = currReachable
			}
		}
	}
	return false
}
