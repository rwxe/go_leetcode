package leetcode

func SearchRange(nums []int, target int) []int {
	leftEnd, rightEnd := 0, len(nums)-1
	l, r := 0, len(nums)-1
	//左边界
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	if l >= len(nums) || nums[l] != target {

		return []int{-1, -1}
	}
	leftEnd = l
	r = len(nums) - 1
	//右边界
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			l = mid + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	rightEnd = r
	//if nums[r] != target || r < 0 {
	//	return []int{-1, -1}
	//}
	return []int{leftEnd, rightEnd}

}
