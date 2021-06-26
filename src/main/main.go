package main

import (
	"fmt"
	"src/leetcode"
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
	fmt.Println(leetcode.IsOneBitCharacter([]int{1,1,0}))
	fmt.Println(leetcode.IsOneBitCharacter([]int{0,0}))
	fmt.Println(leetcode.IsOneBitCharacter([]int{1,1,1,0}))
	fmt.Println(leetcode.IsOneBitCharacter([]int{1,1,1,1,1}))

}
