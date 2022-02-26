package leetcode

func search81(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		} else if nums[l] == nums[mid] && nums[mid] == nums[r] {
			//无法判断左右
			r--
			l++
		} else if nums[l] <= nums[mid] {
			//左有序、只有两个数
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			//右有序
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
