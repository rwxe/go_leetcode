package leetcode

func bt77(nums, path []int, pos, length int, result *[][]int) {
	if len(path) >= length {
		*result = append(*result, append([]int(nil), path...))
		return
	}
	for i := pos; i < len(nums); i++ {
		path = append(path, nums[i])
		bt77(nums, path, i+1, length, result)
		path = path[:len(path)-1]
	}

}

func Combine(n int, k int) [][]int {
	result := make([][]int, 0)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	bt77(nums, []int{}, 0, k, &result)
	return result

}
