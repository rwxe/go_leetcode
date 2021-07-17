package leetcode

// 快排模板
func QuickSort(nums []int,leftEnd,rightEnd int){
	if leftEnd<rightEnd{
		//mid:=hoarePartition(nums,leftEnd,rightEnd)
		mid:=ioaPartition(nums,leftEnd,rightEnd)
		QuickSort(nums,leftEnd,mid-1)
		QuickSort(nums,mid+1,rightEnd)
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

