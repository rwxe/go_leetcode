package leetcode

func nextGreaterElements(nums []int) []int {
	mStack := make([]int, 0)
	result := make([]int, len(nums))
	for i := len(nums)*2 - 1; i >= 0; i-- {
		theIndex := i % len(nums)
		for len(mStack) != 0 && mStack[len(mStack)-1] <= nums[theIndex] {
			mStack = mStack[:len(mStack)-1]
		}
		if len(mStack) == 0 {
			result[theIndex] = -1
		} else {
			result[theIndex] = mStack[len(mStack)-1]
		}
		mStack = append(mStack, nums[theIndex])
	}
	return result
}
