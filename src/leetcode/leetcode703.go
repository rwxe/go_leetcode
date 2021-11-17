package leetcode

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	sort.IntSlice
	k int
}

func (kl *KthLargest) Push(x interface{}) {
	kl.IntSlice = append(kl.IntSlice, x.(int))
}
func (kl *KthLargest) Pop() interface{} {
	old := kl.IntSlice
	n := len(old)
	x := old[n-1]
	kl.IntSlice = old[0 : n-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, v := range nums {
		kl.Add(v)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
