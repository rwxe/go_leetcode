package leetcode

import (
	"container/heap"
	"sort"
)


type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
// 为了实现大根堆，Less在大于时返回小于
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}



func GetLeastNumbers0(arr []int, k int) []int {
	minSeq:=make([]int,0)
	sort.Ints(arr)
	count:=0
	for _,v:=range arr{
		if count< k{
				minSeq=append(minSeq,v)
				count++
		}
	}
	return minSeq

}

func min(arr []int,start int)(int,int){
	min:=arr[start]
	index:=start
	for i:=start;i<len(arr);i++{
		if arr[i]<min{
			min=arr[i]
			index=i
		}
	}
	return index,min
}

// 局部排序
func GetLeastNumbers1(arr []int, k int) []int {
	for i:=0;i<k;i++{
		j,_:=min(arr,i)
		arr[i],arr[j]=arr[j],arr[i]
	}
	return arr[:k]

}

// 大根堆
func GetLeastNumbers2(arr []int, k int) []int {
	h := make(IntHeap,k)
	hp:=&h
	copy(h ,IntHeap(arr[:k+1]))
	heap.Init(hp)
	for i:=k;i<len(arr);i++{
		if arr[i]<h[0]{
			heap.Pop(hp)
			heap.Push(hp,arr[i])
		}
	}
	return h
	
}
