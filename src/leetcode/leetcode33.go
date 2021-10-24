package leetcode

func Search33(nums []int, target int) int {

	//return searcher33R(nums, 0, len(nums)-1, target)
	return searcher33NR(nums, target)
}

func searcher33NR(nums []int, target int) int {
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

func searcher33R(nums []int, left, right, target int) int {

	if left > right {
		return -1
	}
	mid := (left + right) / 2
	midValue := nums[mid]
	leftValue := nums[left]
	rightValue := nums[right]

	if midValue == target {
		return mid

	} else if leftValue <= midValue {

		if leftValue <= target && target <= midValue {
			return searcher33R(nums, left, mid-1, target)

		} else {
			return searcher33R(nums, mid+1, right, target)

		}

	} else {

		if midValue <= target && target <= rightValue {

			return searcher33R(nums, mid+1, right, target)

		} else {
			return searcher33R(nums, left, mid-1, target)

		}

	}

}
