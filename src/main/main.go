package main

import (
	"fmt"
	"go_leetcode/src/algo"
	"go_leetcode/src/leetcode"
)

type set map[int]bool
func (s *set)Add(n int){
	(*s)[n]=true
}

func main() {
	t1:=&algo.TreeNode{Val:1}
	t1.Left=&algo.TreeNode{Val:2}
	t1.Right=&algo.TreeNode{Val:3}
	t2:=&algo.TreeNode{Val:1}
	t2.Left=&algo.TreeNode{Val:2}
	fmt.Println(leetcode.DiameterOfBinaryTree(t2))
	fmt.Println(leetcode.DiameterOfBinaryTree(t1))
	s1:=set{}
	s2:=set{}
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2.Add(2)
	s2.Add(1)
	s2.Add(3)
	s3:=map[int]bool{1:true,2:true,3:true}
	s4:=map[int]bool{2:true,3:true,1:true}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s3==s4)


}
