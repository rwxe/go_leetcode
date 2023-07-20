package leetcode

func maxSubarraySumCircular(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	maxSum := maxSubArray_918(nums)
	minSum := minSubArray_918(nums)
	//fmt.Println(totalSum, maxSum, minSum)
	if totalSum-minSum > maxSum && maxSum > 0 {
		return totalSum - minSum
	} else {
		return maxSum
	}

}

func maxSubArray_918(nums []int) int {
	pre, maxSum := nums[0], nums[0]
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if pre+v > v {
			pre += v
		} else {
			pre = v
		}
		if pre > maxSum {
			maxSum = pre
		}
	}
	return maxSum
}
func minSubArray_918(nums []int) int {
	pre, minSum := nums[0], nums[0]
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if pre+v < v {
			pre += v
		} else {
			pre = v
		}
		if pre < minSum {
			minSum = pre
		}
	}
	return minSum
}
