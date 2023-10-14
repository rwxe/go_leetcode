package main

import (
	"fmt"
	"src/leetcode"
)

func main() {
	fmt.Println(leetcode.GetWordsInLongestSubsequence2(3, []string{"bab", "dab", "cab"}, []int{1, 2, 2}))
	fmt.Println(leetcode.GetWordsInLongestSubsequence2(4, []string{"a", "b", "c", "d"}, []int{1, 2, 3, 4}))
	fmt.Println(leetcode.GetWordsInLongestSubsequence2(3, []string{"bdb", "aaa", "ada"}, []int{2, 1, 3}))
}
