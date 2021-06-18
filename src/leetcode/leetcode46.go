package leetcode

func bt46(list, nums []int, length int, result *[][]int) {
	if len(list) >= length {
		*result = append(*result, append([]int(nil), list...))
		return
	}
	for i := 0; i < len(nums); i++ {
		list = append(list, nums[i])
		newNums := append([]int(nil), nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		bt46(list, newNums, length, result)
		list = list[:len(list)-1]
	}
}

func in46(list []int, elem int) bool {
	for _, v := range list {
		if v == elem {
			return true
		}
	}
	return false
}

func bt46_2(nums []int, path []int, result *[][]int) {
	if len(path) >= len(nums) {
		*result = append(*result, append([]int(nil), path...))
	}
	for i := 0; i < len(nums); i++ {
		if in46(path, nums[i]) {
			continue
		}
		path = append(path, nums[i])
		bt46_2(nums, path, result)
		path = path[:len(path)-1]
	}

}

func Permute(nums []int) [][]int {
	result := make([][]int, 0)
	//bt46([]int{}, nums, n, &result)
	bt46_2(nums, []int{}, &result)
	return result

}
