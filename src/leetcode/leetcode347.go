package leetcode

import (
	"sort"
)

func TopKFrequent347(nums []int, k int) []int {
	numFreqMap := make(map[int]int)
	numsDistinct := make([]int, 0)
	for _, num := range nums {
		numFreqMap[num] += 1
	}
	for k := range numFreqMap {
		numsDistinct = append(numsDistinct, k)
	}
	//fmt.Println("DBEUG", numFreqMap)
	//fmt.Println("DBEUG", numsDistinct)
	sort.Slice(numsDistinct, func(i, j int) bool { return numFreqMap[numsDistinct[i]] > numFreqMap[numsDistinct[j]] })
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, numsDistinct[i])

	}
	return result

}
