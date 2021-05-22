package main

import (
	"fmt"
	"go_leetcode/src/algo"
	"go_leetcode/src/leetcode"
)


func main() {
	t1:=&algo.TreeNode{Val:1}
	t1.Left=&algo.TreeNode{Val:2}
	t1.Right=&algo.TreeNode{Val:3}
	fmt.Println(algo.PostorderTraversal(t1))
	leetcode.InvertTree(t1)
	fmt.Println(algo.PostorderTraversal(t1))


}
