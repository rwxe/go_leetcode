package leetcode

import (
	"container/heap"
	"sort"
)

func MinStoneSum(piles []int, k int) (ans int) {
	h := &hp5839{piles}
	heap.Init(h)
	for ; k > 0; k-- {
		h.IntSlice[0] -= h.IntSlice[0] / 2
		heap.Fix(h, 0)
	}
	for _, v := range h.IntSlice {
		ans += v
	}
	return
}

type hp5839 struct{ sort.IntSlice }

func (h hp5839) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp5839) Push(interface{})     {}
func (hp5839) Pop() (x interface{}) { return }
