package leetcode

import "sort"
func sum40(elems ...int) int {
	sum := 0
	for _, v := range elems {
		sum += v
	}
	return sum
}

func bt40(nums, path []int, pos, target int, result *[][]int) {
	if sum40(path...) > target {
		return
	} else if len(path) > 1 && path[len(path)-1] < path[len(path)-2] {
		return
	} else if sum40(path...) == target {
		*result = append(*result, append([]int(nil), path...))
		return
	}
	for i := pos; i < len(nums); i++ {
		if i>pos && nums[i]==nums[i-1]{
			continue
		}
		path = append(path, nums[i])
		bt40(nums, path, i+1, target, result)
		path = path[:len(path)-1]
	}

}

func CombinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	bt40(candidates, []int{}, 0, target, &result)
	return result

}
