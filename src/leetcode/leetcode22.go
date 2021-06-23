package leetcode



func isValid22(path []rune) bool {
	stack := make([]rune, 0)
	for _, v := range path {
		if v == '(' {
			stack = append(stack, v)
		} else if len(stack)>0{
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0

}
func bt22(nums, path []rune, length int, result *[]string) {
	if len(path) >= length {
		if isValid22(path) {
			*result = append(*result, string(path))
		}
	}
	for i := 0; i < len(nums); i++ {
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		path = append(path, nums[i])
		newNums := append([]rune(nil), nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		bt22(newNums, path, length, result)
		path = path[:len(path)-1]

	}

}

func GenerateParenthesis(n int) []string {
	result := make([]string, 0)
	nums := make([]rune, n*2)
	for i := 0; i < len(nums); i++ {
		if i<len(nums)/2{
			nums[i]='('
		}else{
			nums[i]=')'

		}
	}
	bt22(nums, []rune{}, len(nums), &result)

	return result

}
