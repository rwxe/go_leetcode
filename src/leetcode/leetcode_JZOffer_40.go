package leetcode

import (
	"container/heap"
	"sort"
)

// 大根堆解法
// 前k小用大根堆，前k大用小根堆
type myHeapJZO40 struct {
	sort.IntSlice
}

//因为大根堆，所以覆盖Less方法，返回较大值
func (h myHeapJZO40) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *myHeapJZO40) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *myHeapJZO40) Pop() interface{} {
	x := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return x
}

func getLeastNumbers_0(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	heapArr := make([]int, k)
	copy(heapArr, arr[:k])
	// 重要，取指针
	h := &myHeapJZO40{IntSlice: heapArr}
	heap.Init(h)
	for i := k; i < len(arr); i++ {
		if x := arr[i]; x < h.IntSlice[0] {
			heap.Pop(h)
			heap.Push(h, x)
		}
	}
	return h.IntSlice

}

// 快排
func getLeastNumbers_1(arr []int, k int) []int {
	// partition 返回第k-l+1个小的数
	partition := func(leftEnd, rightEnd int) int {
		pivot := arr[rightEnd]
		l, r := leftEnd-1, leftEnd
		for ; r < rightEnd; r++ {
			if arr[r] <= pivot {
				l++
				arr[l], arr[r] = arr[r], arr[l]
			}
		}
		arr[l+1], arr[rightEnd] = arr[rightEnd], arr[l+1]
		return l + 1
	}
	partition(0, 1)
	//mid==k，得到答案，return arr[:mid]
	//mid<k，寻找右侧，p(mid+1,r,k)
	//TODO
	return []int{}
}
