package main

import (
	"go_leetcode/src/leetcode"
	"testing"
)

func BenchmarkMain0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.Permute0([]int{1,2,3,4,5})
	}
}

func BenchmarkMain1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.Permute1([]int{1,2,3,4,5})
	}
}

func BenchmarkMain2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.Permute2([]int{1,2,3,4,5})
	}
}

