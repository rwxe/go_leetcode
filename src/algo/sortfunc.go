package algo

import (
	"math/rand"
)

//var a =[]int{9,3,41,7,98,621,1,4}

// 快排模板
func QuickSortR(nums []int, leftEnd, rightEnd int) {
	if leftEnd < rightEnd {
		mid := IOAPartition(nums, leftEnd, rightEnd)
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
		stack = stack[:len(stack)-1]
		l := frame[0]
		r := frame[1]
		mid := IOAPartition2(nums, l, r)
		if mid > l {
			stack = append(stack, []int{l, mid - 1})
		}
		if mid < r {
			stack = append(stack, []int{mid + 1, r})
		}

	}
}

// 霍尔分区，采用最左边的元素作为基准
func HoarePartition(nums []int, leftEnd, rightEnd int) int {
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
func IOAPartition(nums []int, leftEnd, rightEnd int) int {
	// leetcode 超时警告
	pivot := nums[rightEnd]
	l := leftEnd - 1
	r := leftEnd
	for ; r < rightEnd; r++ {
		if nums[r] <= pivot {
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[rightEnd], nums[l+1] = nums[l+1], nums[rightEnd]
	return l + 1
}

// 算法导论分区优化1，以最右为基准
// 随机选取基准
func IOAPartition1(nums []int, leftEnd, rightEnd int) int {
	// rand.Seed(time.Now().Unix())
	random := leftEnd + rand.Intn(rightEnd-leftEnd+1)
	nums[rightEnd], nums[random] = nums[random], nums[rightEnd]
	pivot := nums[rightEnd]
	l := leftEnd - 1
	r := leftEnd
	for ; r < rightEnd; r++ {
		if nums[r] <= pivot {
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[rightEnd], nums[l+1] = nums[l+1], nums[rightEnd]
	return l + 1
}

// 算法导论分区优化2，以最右为基准
// 选择三个求中位数
func IOAPartition2(nums []int, leftEnd, rightEnd int) int {
	median := rightEnd
	if len(nums) > 2 {
		half := (leftEnd + rightEnd) / 2
		if nums[leftEnd] <= nums[half] && nums[half] <= nums[rightEnd] {
			median = half
		}
		if nums[leftEnd] <= nums[rightEnd] && nums[rightEnd] <= nums[half] {
			median = rightEnd
		}
		if nums[half] <= nums[leftEnd] && nums[leftEnd] <= nums[rightEnd] {
			median = leftEnd
		}
		if nums[half] <= nums[rightEnd] && nums[rightEnd] <= nums[leftEnd] {
			median = rightEnd
		}
		if nums[rightEnd] <= nums[half] && nums[half] <= nums[leftEnd] {
			median = half
		}
		if nums[rightEnd] <= nums[leftEnd] && nums[leftEnd] <= nums[half] {
			median = leftEnd
		}

	}
	nums[rightEnd], nums[median] = nums[median], nums[rightEnd]
	pivot := nums[rightEnd]
	l := leftEnd - 1
	r := leftEnd
	for ; r < rightEnd; r++ {
		if nums[r] <= pivot {
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[rightEnd], nums[l+1] = nums[l+1], nums[rightEnd]
	return l + 1
}

//归并排序，递归
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	leftNums := MergeSort(arr[:mid])
	rightNums := MergeSort(arr[mid:])
	return merge(leftNums, rightNums)
}

func merge(leftNums, rightNums []int) []int {
	l, r := 0, 0
	result := make([]int, 0)
	for l < len(leftNums) && r < len(rightNums) {
		if leftNums[l] < rightNums[r] {
			result = append(result, leftNums[l])
			l++
		} else {
			result = append(result, rightNums[r])
			r++
		}
	}
	result = append(result, leftNums[l:]...)
	result = append(result, rightNums[r:]...)
	return result
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

}
