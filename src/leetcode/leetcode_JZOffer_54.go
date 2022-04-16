package leetcode

import (
	"container/heap"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type hp_JZO54 struct {
	sort.IntSlice
}

func (hp *hp_JZO54) Pop() interface{} {
	x := hp.IntSlice[len(hp.IntSlice)-1]
	hp.IntSlice = hp.IntSlice[:len(hp.IntSlice)-1]
	return x
}
func (hp *hp_JZO54) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}
func kthLargest(root *TreeNode, k int) int {
	nums := make([]int, k)
	hp := &hp_JZO54{nums}
	heap.Init(hp)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			if node.Val > hp.IntSlice[0] {
				heap.Pop(hp)
				heap.Push(hp, node.Val)
			}
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return nums[0]
}
