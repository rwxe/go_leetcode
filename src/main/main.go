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
	// fmt.Println(leetcode.MinStoneSum([]int{4,3,6,7},3))
	// fmt.Println(leetcode.MinStoneSum([]int{4,3,6,7},5))
	// fmt.Println(leetcode.MinStoneSum([]int{1391,5916},3))
	//fmt.Println(leetcode.MinimumTimeRequired([]int{3,2,3},3))
	//fmt.Println(leetcode.MinimumTimeRequired([]int{1,2,4,7,8},2))
	fmt.Println(leetcode.SearcherNR([]int{3,4,5,1,2},-1<<31))
}
