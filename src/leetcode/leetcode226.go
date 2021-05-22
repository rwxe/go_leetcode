package leetcode

import "go_leetcode/src/algo"

type TreeNode=algo.TreeNode

func preorderTurn(root *TreeNode){
	if root==nil{return}
	root.Left,root.Right=root.Right,root.Left
	preorderTurn(root.Left)
	preorderTurn(root.Right)
}

func InvertTree(root *TreeNode) *TreeNode {
	returnRoot:=root
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			root.Left,root.Right=root.Right,root.Left
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return returnRoot


}
