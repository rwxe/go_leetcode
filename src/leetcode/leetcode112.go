package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs112(root, 0, targetSum)
}

func dfs112(parentNode *TreeNode, currentSum, targetSum int) bool {
	if parentNode != nil {
		currentSum += parentNode.Val
	} else {
		return false
	}
	if currentSum == targetSum && parentNode.Left == nil && parentNode.Right == nil {
		return true
	}
	return dfs112(parentNode.Left, currentSum, targetSum) ||
		dfs112(parentNode.Right, currentSum, targetSum)
}
