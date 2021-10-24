package leetcode

func Rob(nums []int) int {
	memory := make([]int, len(nums))
	for i := range memory {
		memory[i] = -1
	}
	return dp198_1(nums)
}

func dp198_1(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	w := make([]int, len(nums))
	for i := range w {
		if i == 0 {
			w[i] = nums[0]
		} else if i == 1 {
			w[i] = max(nums[1], nums[0])
		} else if i > 1 {
			w[i] = max(w[i-2]+nums[i], w[i-1])
		}
	}
	return w[len(w)-1]
}

func dp198_2(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dp0, dp1, dpResult := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp0 = nums[0]
			dpResult = dp0
		} else if i == 1 {
			dp1 = max(nums[1], nums[0])
			dpResult = dp1
		} else {
			dpResult = max(dp0+nums[i], dp1)
			dp0, dp1 = dp1, dpResult
		}
	}
	return dpResult
}

func dp198_3(nums []int, start int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	if start >= len(nums) {
		return 0
	}
	return max(nums[start]+dp198_3(nums, start+2), dp198_3(nums, start+1))

}
func dp198_4(nums []int, start int, memory []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	if start >= len(nums) {
		return 0
	}
	if memory[start] != -1 {
		return memory[start]
	}
	result := max(nums[start]+dp198_4(nums, start+2, memory), dp198_4(nums, start+1, memory))
	memory[start] = result
	return result

}
