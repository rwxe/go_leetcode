package algo

import _ "fmt"

// 递归前序遍历
func PreorderTraversalR(root *TreeNode, result *[]int) []int {
	if root != nil {
		// 根，左，右
		*result = append(*result, root.Val)
		PreorderTraversalR(root.Left, result)
		PreorderTraversalR(root.Right, result)
	}
	return *result
}

// 递归中序遍历
func InorderTraversalR(root *TreeNode, result *[]int) []int {
	if root != nil {
		// 左，根，右
		InorderTraversalR(root.Left, result)
		*result = append(*result, root.Val)
		InorderTraversalR(root.Right, result)
	}
	return *result
}

// 递归后序遍历
func PostorderTraversalR(root *TreeNode, result *[]int) []int {
	if root != nil {
		// 左，右，根
		*result = append(*result, root.Val)
		PostorderTraversalR(root.Left, result)
		PostorderTraversalR(root.Right, result)
	}
	return *result
}

// 非递归前序遍历
func PreorderTraversalNR(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 非递归中序遍历
func InorderTraversalNR(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 一路向左
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

// 非递归后序遍历
func PostorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}
//从上到下DFS
func PreorderTraversalDFS(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

// V1：深度遍历，结果指针作为参数传入到函数内部
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// 分治法DFS
func PreorderTraversalDAC(root *TreeNode) []int {
	result := make([]int, 0)
	// 返回条件(null & leaf)
	if root == nil {
		return result
	}
	// 分治(Divide)
	left := PreorderTraversalDAC(root.Left)
	right := PreorderTraversalDAC(root.Right)
	// 合并结果(Conquer)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}


//层次遍历
func LevelOrder(root *TreeNode) [][]int {
	// 通过上一层的长度确定下一层的元素
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		level := make([]int, 0)
		// 记录当前层有多少元素（遍历当前层，再添加下一层）
		levelWidth := len(queue)
		for i := 0; i < levelWidth; i++ {
			// 出队列
			levelNode := queue[0]
			queue = queue[1:]
			level = append(level, levelNode.Val)
			if levelNode.Left != nil {
				queue = append(queue, levelNode.Left)
			}
			if levelNode.Right != nil {
				queue = append(queue, levelNode.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *TreeNode, ch chan int) {
	if t == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || t != nil {
		for t != nil {
			stack = append(stack, t)
			t = t.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ch <- node.Val
		t = node.Right
	}
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *TreeNode) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 && ok2 {
			if v1 != v2 {
				return false
			}
		} else if ok1 || ok2 {
			return false
		} else {
			break
		}
	}
	return true
}
