package leetcode

func SumIndicesWithKSetBits(nums []int, k int) int {
	countBinary1 := func(x int) (count int) {
		for x != 0 {
			x &= x - 1
			count++
		}
		return
	}
	if k == 0 {
		return nums[0]
	}
	result := 0
	for i, v := range nums {
		if k == countBinary1(i) {
			result += v
		}
	}
	return result

}
