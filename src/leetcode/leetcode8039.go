package leetcode

func MinimumRightShifts(nums []int) int {
	n := len(nums)
	count := 0
	for start := n; start >= 0; start-- {
		before := nums[start%n]
		for i := start; i < start+n; i++ {
			j := i % n
			if nums[j] < before {
				count++
				goto Restart
			}
			before = nums[j]
		}
		return count
	Restart:
	}
	return -1
}
