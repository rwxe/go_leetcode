package leetcode

func FindShortestSubArray(nums []int) int {
	countMap := make(map[int]int)
	degree := 0
	for _, v := range nums {
		countMap[v]++
	}
	for _, v := range countMap {
		if v > degree {
			degree = v
		}
	}
	modeList := make([]int, 0)
	for k, v := range countMap {
		if v == degree {
			modeList = append(modeList, k)
		}
	}
	minLen := 1 << 31
	for _, m := range modeList {
		start, end := -1, -1
		for i, v := range nums {
			if v == m {
				if start == -1 {
					start = i
				}
				if start != -1 {
					end = i
				}
			}
		}
		currLen := end - start + 1
		if currLen < minLen {
			minLen = currLen
		}
	}
	return minLen
}
func FindShortestSubArray2(nums []int) int {
	type entry struct{ count, left, right int }
	minLen := 0
	maxCount := 0
	numMap := make(map[int]entry)

	for i, v := range nums {
		if e, ok := numMap[v]; ok {
			e.count++
			e.right = i
			numMap[v] = e
		} else {
			numMap[v] = entry{1, i, i}
		}
	}
	for _, e := range numMap {
		if e.count > maxCount {
			maxCount = e.count
			minLen = e.right - e.left + 1
		} else if e.count == maxCount {
			if (e.right - e.left + 1) < minLen {
				minLen = e.right - e.left + 1
			}
		}

	}
	return minLen
}
