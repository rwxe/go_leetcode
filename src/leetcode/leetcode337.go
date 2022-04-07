package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	memory := make(map[*TreeNode]int)
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	var dp func(root *TreeNode) int
	dp = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if _, ok := memory[root]; ok {
			return memory[root]
		}
		robNow := root.Val
		if root.Left != nil {
			robNow += dp(root.Left.Left)
			robNow += dp(root.Left.Right)
		}
		if root.Right != nil {
			robNow += dp(root.Right.Left)
			robNow += dp(root.Right.Right)
		}

		notRob := 0
		if root.Left != nil {
			notRob += dp(root.Left)
		}
		if root.Right != nil {
			notRob += dp(root.Right)
		}

		memory[root] = max(robNow, notRob)
		return memory[root]
	}
	return dp(root)
}
