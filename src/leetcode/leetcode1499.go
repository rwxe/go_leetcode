package leetcode

import (
	"container/heap"
	"math"
)

type heap1499 [][2]int                 //y-x,x
func (h heap1499) Len() int            { return len(h) }
func (h heap1499) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h heap1499) Less(i, j int) bool  { return h[i][0] > h[j][0] } //最大堆
func (h *heap1499) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *heap1499) Pop() interface{}   { x := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return x }
func (h *heap1499) Top() [2]int        { return (*h)[0] }
func (h *heap1499) Bottom() [2]int     { return (*h)[len(*h)-1] }
func (h *heap1499) PopFront() [2]int   { x := (*h)[0]; (*h) = (*h)[1:]; return x }

//堆，优先队列
func findMaxValueOfEquation(points [][]int, k int) int {
	var maxVal int = math.MinInt64
	h := make(heap1499, 0, len(points))
	for _, p := range points {
		x, y := p[0], p[1]
		if h.Len() > 0 {
			for h.Len() > 0 && x-h.Top()[1] > k {
				heap.Pop(&h)
			}
			if h.Len() > 0 {
				//yi-xi+xj+yj
				if val := x + y + h.Top()[0]; val > maxVal {
					maxVal = val
				}
			}
		}
		heap.Push(&h, [2]int{y - x, x})
	}
	return maxVal
}

//单调队列，双端队列
func findMaxValueOfEquation_1(points [][]int, k int) int {
	var maxVal int = math.MinInt64
	//y-x,x
	dq := make(heap1499, 0, k+1) //队列实际上不会超过k+1
	for _, p := range points {
		x, y := p[0], p[1]
		if dq.Len() > 0 {
			for dq.Len() > 0 && x-dq.Top()[1] > k {
				dq.PopFront()
			}
			if dq.Len() > 0 {
				//yi-xi+xj+yj
				if val := x + y + dq.Top()[0]; val > maxVal {
					maxVal = val
				}
			}
			//弹出末尾小于当前yj-xj的yi-xi，让双端队列保持单调递减
			for dq.Len() > 0 && dq.Bottom()[0] < y-x {
				dq.Pop()
			}
		}
		dq = append(dq, [2]int{y - x, x})

	}
	return maxVal
}
