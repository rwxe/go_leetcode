package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mStack := make([]int, 0)
	result1 := make([]int, len(nums1))
	result2 := make([]int, len(nums2))
	nums1ToNums2 := make(map[int]int)

	for i, v := range nums2 {
		nums1ToNums2[v] = i
	}

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(mStack) != 0 && mStack[len(mStack)-1] <= nums2[i] {
			mStack = mStack[:len(mStack)-1]
		}
		if len(mStack) == 0 {
			result2[i] = -1
		} else {
			result2[i] = mStack[len(mStack)-1]
		}
		mStack = append(mStack, nums2[i])
	}

	for i, v := range nums1 {
		result1[i] = result2[nums1ToNums2[v]]
	}
	return result1

}
