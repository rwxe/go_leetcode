package leetcode

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	reverseFlag := false
	for len(queue) > 0 {
		levelWidth := len(queue)
		level := make([]*TreeNode, 0, levelWidth)

		for i := 0; i < levelWidth; i++ {
			levelNode := queue[0]
			queue = queue[1:]
			level = append(level, levelNode)
			if levelNode.Left != nil {
				queue = append(queue, levelNode.Left)
			}
			if levelNode.Right != nil {
				queue = append(queue, levelNode.Right)
			}
		}
		if reverseFlag {
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i].Val, level[j].Val = level[j].Val, level[i].Val
			}
		}
		reverseFlag = !reverseFlag

	}
	return root

}
