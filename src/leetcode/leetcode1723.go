package leetcode


func max1723(nums []int) int {
	result := 0
	for _, v := range nums {
		if v > result {
			result = v
		}
	}
	return result
}
func min1723(nums []int) int {
	result := 1 << 31
	for _, v := range nums {
		if v < result {
			result = v
		}
	}
	return result
}

func checker1723(nums []int) bool {
	for _, v := range nums {
		if v == 0 {
			return false
		}
	}
	return true
}

func bt1723(workers, jobs []int, pos int, result *[]int) {
	if len(jobs) == 0 {
		if checker1723(workers) {
			*result = append(*result, max1723(workers))
		}
		return
	}

	job := jobs[0]
	for i := pos; i < len(workers); i++ {
		workers[i] += job
		bt1723(workers, jobs[1:], 0, result)
		workers[i] -= job
	}
}


func MinimumTimeRequired(jobs []int, k int) int {
	result := 1 << 31
	return result
}
