package leetcode



func sum39(elems ...int) int {
	sum := 0
	for _, v := range elems {
		sum += v
	}
	return sum
}

func bt39(nums, path []int, pos, target int, result *[][]int) {
	if sum39(path...) > target {
		return
	} else if len(path) > 1 && path[len(path)-1] < path[len(path)-2] {
		return
	} else if sum39(path...) == target {
		*result = append(*result, append([]int(nil), path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		bt39(nums, path, i+1, target, result)
		path = path[:len(path)-1]
	}

}

func CombinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	bt39(candidates, []int{}, 0, target, &result)
	return result

}
