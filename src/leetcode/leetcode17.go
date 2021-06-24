package leetcode

func bt17(numPad []string, path []byte, digits string, pos int, result *[]string) {
	if len(path) >= len(digits) {
		if len(digits)>0{
		*result = append(*result, string(path))

		}
		return
	}
	for i := 0; i < len(numPad[(digits[pos]-'0')]); i++ {
		path = append(path, numPad[(digits[pos] - '0')][i])
		bt17(numPad, path, digits, pos+1, result)
		path = path[:len(path)-1]
	}

}

func LetterCombinations(digits string) []string {
	numPad := []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	result := make([]string, 0)
	bt17(numPad, []byte{}, digits, 0, &result)
	return result

}
