package leetcode

func singleNonDuplicate_2(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		// x在mid左边时,重复元素开始为奇数
		// x在mid右边时,重复元素开始为偶数
		// mid为奇数时,mid^1==mid-1
		// mid为偶数时,mid^1==mid+1
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
func singleNonDuplicate_1(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		// x在mid左边时,重复元素开始为奇数
		// x在mid右边时,重复元素开始为偶数
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	return nums[l]
}
func singleNonDuplicate(nums []int) int {
	sum := 0
	operator := 1
	for i := range nums {
		sum += operator * nums[i]
		operator = -operator
		if i != 0 && operator == 1 && sum != 0 {
			return nums[i-1]
		}
	}
	if sum != 0 {
		return nums[len(nums)-1]
	} else {
		return nums[0]

	}
}
