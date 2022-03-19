package main

import (
	"src/leetcode"
	"testing"
)

func Benchmark0(b *testing.B) {
	prices := []int{1, 3, 7, 5, 10, 3}
	for i := 0; i < b.N; i++ {
		leetcode.MaxProfit714_0(prices, 2)
	}
}
func Benchmark1(b *testing.B) {
	prices := []int{1, 3, 7, 5, 10, 3}
	for i := 0; i < b.N; i++ {
		leetcode.MaxProfit714_1(prices, 2)
	}
}
