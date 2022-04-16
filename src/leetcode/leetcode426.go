package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

func treeToDoublyList_2(root *TreeNode) *TreeNode {
	//Left Prev
	//Right Next
	if root == nil {
		return nil
	}
	var haedNode, prevNode *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			if prevNode != nil {
				node.Left = prevNode
				prevNode.Right = node
			} else {
				haedNode = node
			}
			prevNode = node
			dfs(node.Right)
		}
	}
	dfs(root)
	haedNode.Left = prevNode
	prevNode.Right = haedNode
	return haedNode

}
func treeToDoublyList(root *TreeNode) *TreeNode {
	//Left Prev
	//Right Next
	if root == nil {
		return nil
	}
	nodes := make([]*TreeNode, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nodes = append(nodes, root)
		dfs(root.Right)
	}
	dfs(root)
	for i := 0; i < len(nodes); i++ {
		//Prev
		if i == 0 {
			nodes[i].Left = nodes[len(nodes)-1]
		} else {
			nodes[i].Left = nodes[i-1]
		}
		//Next
		if i == 0 && len(nodes) == 1 {
			nodes[i].Right = nodes[i]
		} else if i == len(nodes)-1 {
			nodes[i].Right = nodes[0]
		} else {
			nodes[i].Right = nodes[i+1]
		}
	}

	return nodes[0]

}
