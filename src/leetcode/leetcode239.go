package leetcode

import "container/heap"

type heap239 [][2]int

func (h heap239) Len() int            { return len(h) }
func (h heap239) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h heap239) Less(i, j int) bool  { return h[i][0] > h[j][0] } //最大堆
func (h *heap239) Pop() interface{}   { x := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return x }
func (h *heap239) Push(x interface{}) { (*h) = append((*h), x.([2]int)) }
func (h *heap239) Top() [2]int        { return (*h)[0] }
func (h *heap239) Bottom() [2]int     { return (*h)[len(*h)-1] }
func (h *heap239) PopFront() [2]int   { x := (*h)[0]; (*h) = (*h)[1:]; return x }
func (h *heap239) PopBack() [2]int    { x := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return x }

func maxSlidingWindow_0(nums []int, k int) []int {
	h := make(heap239, 0, len(nums))
	for i := 0; i < k; i++ {
		heap.Push(&h, [2]int{nums[i], i})
	}

	result := make([]int, 0, len(nums)+1-k)
	result = append(result, h.Top()[0])

	for i := k; i < len(nums); i++ {
		heap.Push(&h, [2]int{nums[i], i})
		for h.Len() > 0 && h.Top()[1] <= i-k {
			heap.Pop(&h)
		}
		result = append(result, h.Top()[0])
	}
	return result
}

//单调队列
func maxSlidingWindow(nums []int, k int) []int {
	h := make(heap239, 0, k+1)
	result := make([]int, 0, len(nums)+1-k)
	for i := 0; i < len(nums); i++ {
		for h.Len() > 0 && h.Bottom()[0] < nums[i] {
			h.PopBack()
		}
		h.Push([2]int{nums[i], i})
		for h.Len() > 0 && h.Top()[1] <= i-k {
			h.PopFront()
		}
		if i >= k-1 {
			result = append(result, h.Top()[0])
		}

	}

	return result
}
