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

func Subsets0(nums []int) [][]int {
	result := make([][]int, 0)
	bt78(nums, 0, []int{}, &result)
	return result

}

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
func Subsets2(nums []int) [][]int {
	result := make([][]int, 1)

	for i := 0; i < len(nums); i++ {
		resultLen:=len(result)
		for j := 0; j < resultLen; j++ {
			list := make([]int, len(result[j]))
			copy(list, result[j])
			list = append(list, nums[i])
			result = append(result, list)
		}
	}
	return result

}
