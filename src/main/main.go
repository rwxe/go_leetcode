package main

import (
	"fmt"
	"go_leetcode/src/leetcode"
)

func main() {
	//  	t1:=&algo.TreeNode{Val:1}
	//  	t1.Left=&algo.TreeNode{Val:2}
	//  	t1.Right=&algo.TreeNode{Val:3}
	//  	t1.Left.Left=&algo.TreeNode{Val:4}
	//  	t1.Left.Right=&algo.TreeNode{Val:5}
	//  	t1.Right.Left=&algo.TreeNode{Val:6}
	//  	t1.Right.Right=&algo.TreeNode{Val:7}
	//	fmt.Println(leetcode.LevelOrderJZ32(t1))
	fmt.Println(leetcode.Permute([]int{1, 2, 3}))
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1,len(slice1),cap(slice1))
	newSlice := append(slice1[:2], slice1[3:]...)
	n2 := append(slice1,1,2,3,4,54,33)
	fmt.Println(newSlice,len(newSlice),cap(newSlice))
	fmt.Println(slice1,len(slice1),cap(slice1))

	fmt.Printf("n2,%v%T%p\n", n2, n2, n2)
	fmt.Printf("slice1,%v%T%p\n", slice1, slice1, slice1)
	fmt.Printf("slice1,%v%T%p\n", slice1, slice1, slice1)
	fmt.Printf("newSlice,%v%T%p\n", newSlice, newSlice, newSlice)
}
