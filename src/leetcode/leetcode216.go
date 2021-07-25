package leetcode

func sum216(elems ...int) int {
	sum := 0
	for _, v := range elems {
		sum += v
	}
	return sum
}

func bt216(nums, path []int, pos, length, target int, result *[][]int) {
	if sum216(path...) > target {
		return
	} else if len(path) > length {
		return
	} else if sum216(path...) == target && len(path) == length {
		*result = append(*result, append([]int(nil), path...))
		return
	}
	for i := pos; i < len(nums); i++ {
		path = append(path, nums[i])
		bt216(nums, path, i+1, length, target, result)
		path = path[:len(path)-1]
	}

}
func bt216_2(nums, path []int, pos, length, target int, result *[][]int) {
	if target < 0 {
		return
	} else if len(path) > length {
		return
	} else if target == 0 && len(path) == length {
		*result = append(*result, append([]int(nil), path...))
		return
	}
	for i := pos; i < len(nums); i++ {
		path = append(path, nums[i])
		bt216_2(nums, path, i+1, length, target-nums[i], result)
		path = path[:len(path)-1]
	}
}
func CombinationSum3(k, n int) [][]int {
	result := make([][]int, 0)
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bt216(candidates, []int{}, 0, k, n, &result)
	return result

}
