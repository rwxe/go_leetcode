package leetcode

func singleNumber(nums []int) []int {
	// x == res1^res2
	x := 0
	for _, v := range nums {
		x ^= v
	}
	//lowbit
	lowbit := x & -x
	//x0 lowbit位为0
	//x1 lowbit位为1
	x0, x1 := 0, 0
	for _, v := range nums {
		if v&lowbit == lowbit {
			x1 ^= v
		} else {
			x0 ^= v
		}
	}

	return []int{x0, x1}
}
