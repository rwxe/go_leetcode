package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func expandBinaryTree(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val != -1 {
			if root.Left != nil {
				newLeft := &TreeNode{Val: -1}
				newLeft.Left = root.Left
				root.Left = newLeft
			}
			if root.Right != nil {
				newRight := &TreeNode{Val: -1}
				newRight.Right = root.Right
				root.Right = newRight
			}
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return root

}
