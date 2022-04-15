package main

import (
	"fmt"
	"src/leetcode"
)

func solution1(s string) []string {
	reusult := make([]string, 0)
	runes := []rune(s)
	var bt func(path []rune, pos int)
	bt = func(path []rune, pos int) {
		reusult = append(reusult, string(path))
		for i := pos; i < len(runes); i++ {
			path = append(path, rune(runes[i]))
			bt(path, i+1)
			path = path[:len(path)-1]
		}
	}
	bt([]rune{}, 0)
	return reusult
}

func solution(s, replace string) string {
	replaceSet := make(map[rune]struct{}, len(replace))
	for _, c := range replace {
		replaceSet[c] = struct{}{}
	}
	result := make([]rune, 0, len(s))
	for _, c := range s {
		if _, ok := replaceSet[c]; !ok {
			result = append(result, c)
		}
	}
	return string(result)
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
	//fmt.Println(leetcode.MinStoneSum([]int{4,3,6,7},3))
	//fmt.Println(leetcode.MinStoneSum([]int{4,3,6,7},5))
	//fmt.Println(leetcode.MinStoneSum([]int{1391,5916},3))
	//fmt.Println(leetcode.MinimumTimeRequired([]int{3,2,3},3))
	//fmt.Println(leetcode.MinimumTimeRequired([]int{1,2,4,7,8},2))
	//fmt.Println(leetcode.MaxDistance([]int{1, 8, 3, 8, 3}))
	//fmt.Println(leetcode.MaxDistance([]int{0, 1}))
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	nums1 := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(leetcode.Trap42_0(nums))
	fmt.Println(leetcode.Trap42_0(nums1))
	fmt.Println(solution1("abc"))
	//fmt.Println(leetcode.MaxProfit123_1(nums))
	//fmt.Println(leetcode.GetAncestors(8, [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}))

}
