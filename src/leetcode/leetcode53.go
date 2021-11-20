package leetcode

import "src/algo"

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
		if currSum+num > num {
			currSum += num
		} else {
			currSum = num
		}
		if maxSum < currSum {
			maxSum = currSum
		}
	}
	return maxSum

}
func MaxSubArrayDP_1(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = algo.MaxInts(dp[i-1]+nums[i], nums[i])
	}
	return algo.MaxInts(dp...)
}

func MaxSubArrayDP_2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if maxSum < nums[i] {
			maxSum = nums[i]
		}
	}
	return maxSum
}

// 使用线段树
func MaxSubArrayST(nums []int) int {
	max := func(items ...int) int {
		maxItem := items[0]
		for _, v := range items {
			if maxItem < v {
				maxItem = v
			}
		}
		return maxItem
	}
	//l,r区间状态
	type segmentStatus struct {
		iSum int // 区间总和
		lSum int // l为起点区间最大子序和
		rSum int // r为终点区间最大子序和
		mSum int // 区间最大子序和
	}

	var getSegmentMaxSubArray func(l, r int) segmentStatus
	var pushUpSegmentMaxSubArray func(lss, rss segmentStatus) segmentStatus

	pushUpSegmentMaxSubArray = func(lss, rss segmentStatus) segmentStatus {
		iSum := lss.iSum + rss.iSum
		lSum := max(lss.lSum, lss.iSum+rss.lSum)
		rSum := max(rss.rSum, rss.iSum+lss.rSum)
		mSum := max(lss.mSum, rss.mSum, lss.rSum+rss.lSum)
		return segmentStatus{iSum, lSum, rSum, mSum}
	}

	getSegmentMaxSubArray = func(l, r int) segmentStatus {
		if l == r {
			return segmentStatus{nums[l], nums[l], nums[l], nums[l]}
		}
		m := (l + r) / 2
		lSub := getSegmentMaxSubArray(l, m)
		rSub := getSegmentMaxSubArray(m+1, r)
		return pushUpSegmentMaxSubArray(lSub, rSub)
	}
	return getSegmentMaxSubArray(0, len(nums)).mSum

}
