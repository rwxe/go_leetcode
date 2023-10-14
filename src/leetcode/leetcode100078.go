package leetcode

func GetWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	seq := make([]int, 0, n)
	pg := -1
	for i, g := range groups {
		if g != pg {
			seq = append(seq, i)
		}
		pg = g
	}
	//fmt.Println(seq)
	result := make([]string, 0, len(seq))
	for _, gi := range seq {
		result = append(result, words[gi])
	}
	return result
}
