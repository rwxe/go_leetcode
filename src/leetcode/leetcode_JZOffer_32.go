package leetcode

func LevelOrderJZ32(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	left := true
	for len(queue) > 0 {
		level := make([]int, 0)
		width := len(queue)
		if left {
			left = false
			for i := 0; i < width; i++ {
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
		} else {
			left = true
			for i := width; i > 0; i-- {
				node := queue[0]
				queue = queue[1:]
				level = append(level, node.Val)
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
				if node.Left != nil {
					queue = append(queue, node.Left)
				}

			}
		}
		res = append(res, level)

	}
	return res
}
