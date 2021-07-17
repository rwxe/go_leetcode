package main

import (
	"src/leetcode"
	"testing"
)

//var nums1 []int = []int{9,8,7,6,5,4,3,20,1}
var nums1 []int = []int{9,8,7,6,5,4,3,2,1}

func BenchmarkMain0(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}


func BenchmarkMain1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.QuickSort(nums1,0,len(nums1)-1)
	}
}
