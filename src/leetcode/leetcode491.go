package leetcode

func bt(nums, path []int, pos int, result *[][]int) {
	if len(path) > 1 {
		*result = append(*result, append([]int(nil), path...))
	}
	used := make(map[int]bool)
	for i := pos; i < len(nums); i++ {
		if used[nums[i]] || (len(path) > 0 && path[len(path)-1] > nums[i]) {
			continue
		}
		path = append(path, nums[i])
		used[nums[i]] = true
		bt(nums, path, i+1, result)
		path = path[:len(path)-1]
	}
}

func FindSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	bt(nums, []int{}, 0, &result)
	return result

}
