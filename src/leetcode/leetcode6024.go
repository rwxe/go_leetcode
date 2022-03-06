package leetcode

func mostFrequent(nums []int, key int) int {
	targets := make([][]int, 0, len(nums)/2)
	targetCount := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			targetCount[nums[i+1]]++
		}
	}
	for k, v := range targetCount {
		targets = append(targets, []int{k, v})
	}
	maxTimes := 0
	mostItem := 0
	for _, v := range targets {
		if v[1] > maxTimes {
			maxTimes = v[1]
			mostItem = v[0]
		}
	}
	return mostItem
}
