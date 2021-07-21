package leetcode

// 快排模板,递归
func QuickSortR(nums []int, leftEnd, rightEnd int) {
	if leftEnd < rightEnd {
		//mid:=hoarePartition(nums,leftEnd,rightEnd)
		mid := ioaPartition(nums, leftEnd, rightEnd)
		QuickSortR(nums, leftEnd, mid-1)
		QuickSortR(nums, mid+1, rightEnd)
	}
}

// 快排模板,非递归
func QuickSortNR(nums []int, leftEnd, rightEnd int) {
	stack := make([][]int, 0)
	stack = append(stack, []int{leftEnd, rightEnd})
	for len(stack) != 0 {
		frame := stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		l := frame[0]
		r := frame[1]
		mid := ioaPartition(nums, l, r)
		if mid > l {
			stack = append(stack, []int{l, mid - 1})
		}
		if mid < r {
			stack = append(stack, []int{mid + 1, r})
		}

	}
}

// 霍尔分区，采用最左边的元素作为基准
func hoarePartition(nums []int, leftEnd, rightEnd int) int {

	pivot := nums[leftEnd]
	l := leftEnd
	r := rightEnd
	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		for l < r && nums[l] <= pivot {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[leftEnd], nums[l] = nums[l], nums[leftEnd]
	return l
}

// 算法导论分区，以最右为基准
func ioaPartition(nums []int, leftEnd, rightEnd int) int {
	// leetcode 超时警告

	pivot := nums[rightEnd]
	l := leftEnd - 1
	r := leftEnd
	for ; r <= rightEnd-1; r++ {
		if nums[r] <= pivot {
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[rightEnd], nums[l+1] = nums[l+1], nums[rightEnd]
	return l + 1
}
