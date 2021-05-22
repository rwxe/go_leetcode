package leetcode


func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	currSum := nums[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if currSum+num > num{
			currSum+=num
		}else{
			currSum=num
		}
		if maxSum<currSum{
			maxSum=currSum
		}
	}
	return maxSum

}

func MaxSubArrayDP(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum:=nums[0]
	for i:=1 ; i < len(nums);i++{
		if nums[i-1]>0{
			nums[i]+=nums[i-1]
		}
		if maxSum<nums[i]{
			maxSum=nums[i]
		}
	}
	return maxSum
	
}
