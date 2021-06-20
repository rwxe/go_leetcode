package main

import (
	"go_leetcode/src/leetcode"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.CountPoints([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}})
	}
}

func BenchmarkMain2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.CountPoints2([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}})
	}
}

