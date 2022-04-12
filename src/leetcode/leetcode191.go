package leetcode

func hammingWeight_2(num uint32) int {
	//lowbit
	count := 0
	for num != 0 {
		count++
		num -= num & -num
	}
	return count
}
func hammingWeight_1(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}
func hammingWeight_0(num uint32) int {
	count := 0
	var p uint32 = 1
	for ; num != 0; num >>= 1 {
		if p&num == 1 {
			count++
		}
	}
	return count
}
