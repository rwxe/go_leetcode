package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	theMap := make(map[int]struct{})
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if _, ok := theMap[k-root.Val]; ok {
			return true
		}
		theMap[root.Val] = struct{}{}
		return dfs(root.Left) || dfs(root.Right)
	}
	return dfs(root)
}
