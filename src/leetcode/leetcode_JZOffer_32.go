package leetcode

func reverseLevel(level []int) {
	for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
		level[i], level[j] = level[j], level[i]

	}

}

func LevelOrderJZ32(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelWidth := len(queue)
		level := make([]int, 0)
		for i := 0; i < levelWidth; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	left := false
	for i := range result {
		if left {
			left = false
			reverseLevel(result[i])

		} else {
			left = true
		}

	}
	return result

}
