package leetcode

import "fmt"

func SubarraySum_2(nums []int, k int) int {
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
func SubarraySum(nums []int, k int) int {
    count:=0
    preSumMap:=make(map[int]int)
    preSumMap[0]=1
    currentPreSum:=0
    for i:=0;i<len(nums);i++{
        currentPreSum+=nums[i]
        target:=currentPreSum-k
        if times,ok:=preSumMap[target];ok{
            count+=times
        }
        preSumMap[currentPreSum]++
    }
    return count
}
