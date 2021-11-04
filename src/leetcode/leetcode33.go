package leetcode

func search(nums []int, target int) int {
	return searcherR(nums, target, 0, len(nums)-1)
	//return searcherNR(nums, target)
}

func searcherNR(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
	}
	return -1

}

func searcherR(nums []int, target, left, right int) int {
	mid := (left + right) / 2
	if left > right {
		return -1
	} else if nums[mid] == target {
		return mid
	} else if nums[left] <= nums[mid] {
		if nums[left] <= target && target <= nums[mid] {
			right = mid - 1
			return searcherR(nums, target, left, right)
		} else {
			left = mid + 1
			return searcherR(nums, target, left, right)
		}

	} else {
		if nums[mid] <= target && target <= nums[right] {
			left = mid + 1
			return searcherR(nums, target, left, right)
		} else {
			right = mid - 1
			return searcherR(nums, target, left, right)
		}

	}
}
