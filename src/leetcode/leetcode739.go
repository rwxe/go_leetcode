package leetcode

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	mStack := make([]int, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(mStack) != 0 && temperatures[mStack[len(mStack)-1]] <= temperatures[i] {
			mStack = mStack[:len(mStack)-1]
		}
		if len(mStack) == 0 {
			result[i] = 0
		} else {
			result[i] = mStack[len(mStack)-1] - i
		}
		mStack = append(mStack, i)
	}
	return result

}
