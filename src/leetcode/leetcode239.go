package leetcode

import "container/heap"

type heap239 [][2]int                 //num,index
func (h heap239) Len() int            { return len(h) }
func (h heap239) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h heap239) Less(i, j int) bool  { return h[i][0] > h[j][0] } //最大堆
func (h *heap239) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *heap239) Pop() interface{}   { x := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return x }
func (h *heap239) Top() [2]int        { return (*h)[0] }
func (h *heap239) Bottom() [2]int     { return (*h)[len(*h)-1] }
func (h *heap239) PopFront() [2]int   { x := (*h)[0]; (*h) = (*h)[1:]; return x }

//优先队列，堆
func maxSlidingWindow(nums []int, k int) []int {
	h := make(heap239, k, len(nums))
	for i := 0; i < k; i++ {
		h = append(h, [2]int{nums[i], i})
	}
	heap.Init(&h)
	result := make([]int, 0, len(nums)-k+1)
	result = append(result, h.Top()[0])
	for i := k; i < len(nums); i++ {
		windowLeft := i - k
		heap.Push(&h, [2]int{nums[i], i})
		for h.Top()[1] <= windowLeft {
			heap.Pop(&h)
		}
		result = append(result, h.Top()[0])
	}
	return result

}

//双端队列，单调队列
func maxSlidingWindow_1(nums []int, k int) []int {
	dq := make(heap239, k, len(nums))
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		windowLeft := i - k
		for dq.Len() > 0 && dq.Bottom()[0] < v {
			dq.Pop()
		}
		dq.Push([2]int{v, i})
		for dq.Top()[1] <= windowLeft {
			dq.PopFront()
		}
		if i >= k-1 {
			result = append(result, dq.Top()[0])
		}
	}
	return result

}
