package leetcode

func isValidMT0809(nums []byte) bool {
	flag := 0
	for _, v := range nums {
		if flag<0{
			return false
		}
		if v == '(' {
			flag++
		} else {
			flag--
		}
	}
	return flag == 0
}
func btMST0809(nums, path []byte, length int, result *[]string) {
	if len(path) == length {
		if isValidMT0809(path) {
			*result = append(*result, string(append([]byte(nil), path...)))
		}
		return
	}
	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		btMST0809(nums, path, length, result)
		path = path[:len(path)-1]
	}
}

func GenerateParenthesis2(n int) []string {
	result := make([]string, 0)
	btMST0809([]byte{'(', ')'}, []byte{}, n*2, &result)
	return result

}
