package leetcode

/**
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 1
	xDF := []int{0, 0}
	yDF := []int{0, 0}

	for len(queue) > 0 {

		levelLen := len(queue)
		for i := 0; i < levelLen; i++ {
			levelNode := queue[0]
			queue = queue[1:]
			if levelNode != nil {
				if levelNode.Left.Val == x {
					xDF[0] = depth
					xDF[1] = levelNode.Val

				}
				if levelNode.Right.Val == x {
					xDF[0] = depth
					xDF[1] = levelNode.Val

				}
				if levelNode.Left.Val == y || levelNode.Right.Val == y {
					yDF[0] = depth
					yDF[1] = levelNode.Val

				}
			}
			if levelNode.Left != nil {
				queue = append(queue, levelNode.Left)
			}
			if levelNode.Right != nil {
				queue = append(queue, levelNode.Right)
			}

		}
		depth += 1

	}
	if xDF[0]!=yDF[0]{
		return false
	}else if xDF[1]==yDF[1]{
		return false
	}else{
		return true
	}
}
