package leetcode

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil || root2 == nil {
		if root1 == nil {
			return root2
		} else if root2 == nil {
			return root1
		}
	}
	queue1:=make([]*TreeNode,0)

}
