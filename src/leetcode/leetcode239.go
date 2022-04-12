package leetcode

import (
	"container/heap"
	"sort"
)

type MyHeap239 struct {
	nums      []int
	indexHeap sort.IntSlice
}

func (mh *MyHeap239) Len() int {
	return len(mh.indexHeap)
}
func (mh *MyHeap239) Swap(i, j int) {
	mh.indexHeap[i], mh.indexHeap[j] = mh.indexHeap[j], mh.indexHeap[i]
}

//大根堆，所以是">
func (mh *MyHeap239) Less(i, j int) bool {
	return mh.nums[mh.indexHeap[i]] > mh.nums[mh.indexHeap[j]]
}

func (mh *MyHeap239) Push(x interface{}) {
	mh.indexHeap = append(mh.indexHeap, x.(int))
}
func (mh *MyHeap239) Pop() interface{} {
	x := mh.indexHeap[len(mh.indexHeap)-1]
	mh.indexHeap = mh.indexHeap[:len(mh.indexHeap)-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	myHeap := &MyHeap239{nums: nums, indexHeap: make([]int, k)}
	for i := 0; i < k; i++ {
		myHeap.indexHeap[i] = i
	}
	heap.Init(myHeap)
	result := make([]int, 0, len(nums)-k+1)
	result = append(result, myHeap.nums[myHeap.indexHeap[0]])
	for l, r := 1, k; r < len(nums); l, r = l+1, r+1 {
		heap.Push(myHeap, r)
		for myHeap.indexHeap[0] < l {
			heap.Pop(myHeap)
		}
		result = append(result, myHeap.nums[myHeap.indexHeap[0]])
	}
	return result
}

func maxSlidingWindow_0(nums []int, k int) []int {
	//超时
	// lpop,rpush
	// rpush>=currMax?currMax=rpush:else if
	// lpop==currMax?currMax=max(nums[l:r+1]):currMax不变
	max := func(items ...int) int {
		theMax := items[0]
		for i := range items {
			if theMax < items[i] {
				theMax = items[i]
			}
		}
		return theMax
	}
	if k >= len(nums) {
		return []int{max(nums...)}
	}
	result := make([]int, 0, len(nums)-k+1)
	l, r := 0, k-1
	currMax := max(nums[l : r+1]...)
	for ; r < len(nums); l, r = l+1, r+1 {
		if l != 0 {
			lpop := nums[l-1]
			rpush := nums[r]
			if rpush >= currMax {
				currMax = rpush
			} else if lpop == currMax {
				currMax = max(nums[l : r+1]...)
			}
		}
		result = append(result, currMax)
	}
	return result
}
