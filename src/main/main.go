package main

import (
	"fmt"
	"go_leetcode/src/leetcode"
)

func main() {
//	  	t1:=&algo.TreeNode{Val:1}
//	  	t1.Left=&algo.TreeNode{Val:2}
//	  	t1.Right=&algo.TreeNode{Val:3}
//	  	t1.Left.Left=&algo.TreeNode{Val:4}
//	  	t1.Left.Right=&algo.TreeNode{Val:5}
//	  	t1.Right.Left=&algo.TreeNode{Val:6}
//	  	t1.Right.Right=&algo.TreeNode{Val:7}
//		fmt.Println(algo.LevelOrder(t1))
//		fmt.Println(algo.PreorderTraversalR(t1,&[]int{}))
//		fmt.Println(algo.InorderTraversalR(t1,&[]int{}))
		fmt.Println(leetcode.PermuteUnique([]int{1,2,1,2,1,2}))
		fmt.Println(leetcode.GenerateParenthesis(3))

}
