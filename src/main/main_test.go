package main

import (
	"src/leetcode"
	"testing"
)

//var nums1 []int = []int{9,8,7,6,5,4,3,2,1}
var nums1 []int = []int{1,2,3,9,4,5,6,7,2,3,5,7,8}
var can []int=[]int{47,36,30,38,22,41,45,35,32,28,46,48,40,24,39,23,42,21,33,43,31,26,27,37,29,34,49,20}
var target=61

func BenchmarkCombinationSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.CombinationSum_1(can,target)
	}
}


func BenchmarkCombinationSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.CombinationSum_2(can,target)
	}
}


func BenchmarkCombinationSum3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.CombinationSum_3(can,target)
	}
}


