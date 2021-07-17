package algo

import (
	"math/rand"
	"time"
)

//var a =[]int{9,3,41,7,98,621,1,4}

// 快排模板
func QuickSort(nums []int, leftEnd, rightEnd int) {
	if leftEnd < rightEnd {
		mid := ioaPartition(nums, leftEnd, rightEnd)
		QuickSort(nums, leftEnd, mid-1)
		QuickSort(nums, mid+1, rightEnd)
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

// 算法导论分区优化1，以最右为基准
// 随机选取基准
func ioaPartition1(nums []int, leftEnd, rightEnd int) int {
	random := leftEnd + rand.Intn(rightEnd-leftEnd+1)
	nums[rightEnd], nums[random] = nums[random], nums[rightEnd]
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

// 算法导论分区优化2，以最右为基准
// 选择三个求中位数
func ioaPartition2(nums []int, leftEnd, rightEnd int) int {
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
	for ; r <= rightEnd-1; r++ {
		if nums[r] <= pivot {
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[rightEnd], nums[l+1] = nums[l+1], nums[rightEnd]
	return l + 1
}

func merge(left, right []int) []int {
	l := 0
	r := 0
	result := make([]int, 0)
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	result := merge(left, right)
	return result
}

func QuickSort0(arr []int, leftEnd, rightEnd int) {
	if leftEnd > rightEnd {
		return
	}
	left := leftEnd
	right := rightEnd
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right -= 1
		}
		if left < right {
			arr[left] = arr[right]
		}
		for left < right && arr[left] <= pivot {
			left += 1
		}
		if left < right {
			arr[right] = arr[left]
		}
		if left >= right {
			arr[left] = pivot
		}
	}
	QuickSort0(arr, leftEnd, right-1)
	QuickSort0(arr, left+1, rightEnd)
}

