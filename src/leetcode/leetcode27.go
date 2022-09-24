package leetcode

func removeElement_2(nums []int, val int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == val {
			nums[l] = nums[r]
			r--
		} else {
			l++
		}
	}
	return l
}

// 交换
func removeElement(nums []int, val int) int {

	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == val {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		} else {
			l++
		}
	}
	return l
}
