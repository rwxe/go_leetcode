package main

import (
	"testing"
)

func Benchmark0(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			sum++
		} else {
			sum--
		}
	}
}
func Benchmark1(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		if i&1 == 0 {
			sum++
		} else {
			sum--
		}
	}
}
