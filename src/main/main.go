package main

import (
	"fmt"
	"src/leetcode"
)

func map1(m map[int]bool) {
	m[2] = true
}

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
	board := [][]byte{
		{'.', '.', 'W', '.', 'B', 'W', 'W', 'B'},
		{'B', 'W', '.', 'W', '.', 'W', 'B', 'B'},
		{'.', 'W', 'B', 'W', 'W', '.', 'W', 'W'},
		{'W', 'W', '.', 'W', '.', '.', 'B', 'B'},
		{'B', 'W', 'B', 'B', 'W', 'W', 'B', '.'},
		{'W', '.', 'W', '.', '.', 'B', 'W', 'W'},
		{'B', '.', 'B', 'B', '.', '.', 'B', 'B'},
		{'.', 'W', '.', 'W', '.', 'W', '.', 'W'}}
	rMove, cMove := 5, 4
	color := byte('W')

	fmt.Println(leetcode.CheckMove(board, rMove, cMove, color))

}
