package leetcode

//快慢指针
func findDuplicate_3(nums []int) int {
	//TODO
	return -1
}

//二进制
func findDuplicate_2(nums []int) int {
	bitMax := 32
	for n := len(nums) - 1; (n>>bitMax)&1 == 0; {
		bitMax--
	}
	result := 0
	for bit := 0; bit <= bitMax; bit++ {
		x, y := 0, 0
		for i, v := range nums {
			if v>>bit&1 == 1 {
				x++
			}
			if i>>bit&1 == 1 {
				y++
			}
		}
		if x > y {
			result |= 1 << bit
		}
	}
	return result
}

//二分
func findDuplicate_1(nums []int) int {
	var result int
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count > mid {
			result = mid
			r = mid - 1
		} else if count <= mid {
			l = mid + 1
		}
	}
	return result
}
