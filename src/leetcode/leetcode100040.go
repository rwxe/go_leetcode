package leetcode

func CountWays(nums []int) int {
	path := make(map[int]struct{}, len(nums))
	result := 0
	bt100040(nums, 0, path, &result)
	return result
}

func bt100040(nums []int, pos int, path map[int]struct{}, result *int) {
	//*result = append(*result, append([]int(nil), path...))
	//check all happy
	allHappy := true
	for i, v := range nums {
		if _, ok := path[i]; ok && len(path) <= v ||
			!ok && len(path) >= v {
			allHappy = false
		}
	}
	if allHappy {
		*result++
	}

	for i := pos; i < len(nums); i++ {
		path[i] = struct{}{}
		bt100040(nums, i+1, path, result)
		delete(path, i)
	}
}
