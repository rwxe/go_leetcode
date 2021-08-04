package leetcode

import "fmt"

func SubarraySum(nums []int, k int) int {
	count := 0
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(prefixSum)-1; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			// nums[j:i]
			fmt.Println(nums[j:i])
			if prefixSum[i]-prefixSum[j] == k {
				count++
			}
		}
	}

	return count

}
func SubarraySum_1(nums []int, k int) int {
	count := 0
	prefixSum := make(map[int]int)
	prefixSum[0]=1
	prefixSumI := 0
	for i := 0; i < len(nums); i++ {
		prefixSumI += nums[i]
		prefixSumJ := prefixSumI - k
		fmt.Println(prefixSumI,prefixSumJ)
		if times, ok := prefixSum[prefixSumJ]; ok {
			count += times
		}
		prefixSum[prefixSumI]++
	}

	return count

}
