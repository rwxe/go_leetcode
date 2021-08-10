package leetcode

func FindMaxLength(nums []int) int {
	maxLength := 0
	prefixSumMap := make(map[int]int)
	prefixSumMap[0] = -1
	currentPreSum := 0
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	for i, v := range nums {
		currentPreSum += v
		if preIndex, ok := prefixSumMap[currentPreSum]; ok {
			if (i - preIndex) > maxLength {
				maxLength = i - preIndex
			}
		} else {
			prefixSumMap[currentPreSum] = i
		}

	}
	return maxLength
}
func FindMaxLength_2(nums []int) int {
	maxLength := 0
	prefixSum := make([]int, len(nums)+1)
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	for i := 0; i < len(prefixSum)-1; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if prefixSum[j]-prefixSum[i] == 0 {
				if j-i > maxLength {
					maxLength = j - i
				}
			}
		}
	}
	return maxLength

}
