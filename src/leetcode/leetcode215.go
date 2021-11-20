package leetcode

import (
	"container/heap"
	"sort"
)

type IntHeap215 struct {
	sort.IntSlice
}

func (ih *IntHeap215) Push(x interface{}) {
	ih.IntSlice = append(ih.IntSlice, x.(int))
}
func (ih *IntHeap215) Pop() interface{} {
	old := ih.IntSlice
	n := len(old)
	x := old[n-1]
	ih.IntSlice = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	ih := IntHeap215{}
	for _, v := range nums {
		heap.Push(&ih, v)
		if ih.Len() > k {
			heap.Pop(&ih)
		}
	}
	return ih.IntSlice[0]

}
