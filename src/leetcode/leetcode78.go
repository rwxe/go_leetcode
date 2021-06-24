package leetcode

func bt78(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		bt78(nums, i+1, list, result)
		list = list[:len(list)-1]
	}

}

func bt78_2(nums []int, pos int, path []int, result *[][]int) {
	*result = append(*result, append([]int(nil), path...))
	for i := pos; i < len(nums); i++ {
		path = append(path, nums[i])
		bt78_2(nums, i+1, path, result)
		path = path[:len(path)-1]
	}

}

// 回溯
func Subsets0(nums []int) [][]int {
	result := make([][]int, 0)
	//bt78(nums, 0, []int{}, &result)
	bt78_2(nums, 0, []int{}, &result)
	return result

}

// 位运算
func Subsets1(nums []int) [][]int {
	seqMax := (1 << len(nums)) - 1
	result := make([][]int, 0)
	for mask := 0; mask <= seqMax; mask++ {
		list := make([]int, 0)
		for i, v := range nums {
			if (mask>>i)&1 > 0 {
				list = append(list, v)
			}

		}
		result = append(result, list)
	}
	return result

}

// BFS
func Subsets2(nums []int) [][]int {
	// 注意初始长度为1，包含空集
	result := make([][]int, 1)

	for i := 0; i < len(nums); i++ {
		resultLen := len(result)
		for j := 0; j < resultLen; j++ {
			list := append([]int(nil), result[j]...)
			list = append(list, nums[i])
			result = append(result, list)
		}
	}
	return result

}

// 回溯2
func Subsets3(nums []int) [][]int {
	result := make([][]int, 0)
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			result = append(result, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return result
}
