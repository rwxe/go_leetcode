package main

import (
	"go_leetcode/src/leetcode"
	"testing"
)
func BenchmarkCP0(b *testing.B){
	for i:=0;i<b.N;i++{
		leetcode.CountPrimes0(500000)
	}
}
func BenchmarkCP1(b *testing.B){
	for i:=0;i<b.N;i++{
		leetcode.CountPrimes1(500000)
	}
}


