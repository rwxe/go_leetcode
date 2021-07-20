package main

import (
	"math/rand"
	"src/algo"
	"testing"
)

//var nums1 []int = []int{9,8,7,6,5,4,3,2,1}
var nums1 []int = []int{1,2,3,9,4,5,6,7,2,3,5,7,8}

func QuickSort0(nums []int, leftEnd, rightEnd int) {
	if leftEnd < rightEnd {
		mid := algo.HoarePartition(nums, leftEnd, rightEnd)
		QuickSort0(nums, leftEnd, mid-1)
		QuickSort0(nums, mid+1, rightEnd)
	}
}
func QuickSort1(nums []int, leftEnd, rightEnd int) {
	if leftEnd < rightEnd {
		mid := algo.IOAPartition(nums, leftEnd, rightEnd)
		QuickSort1(nums, leftEnd, mid-1)
		QuickSort1(nums, mid+1, rightEnd)
	}
}
func QuickSort2(nums []int, leftEnd, rightEnd int) {
	if leftEnd < rightEnd {
		mid := algo.IOAPartition1(nums, leftEnd, rightEnd)
		QuickSort2(nums, leftEnd, mid-1)
		QuickSort2(nums, mid+1, rightEnd)
	}
}
func QuickSort3(nums []int, leftEnd, rightEnd int) {
	if leftEnd < rightEnd {
		mid := algo.IOAPartition2(nums, leftEnd, rightEnd)
		QuickSort3(nums, leftEnd, mid-1)
		QuickSort3(nums, mid+1, rightEnd)
	}
}

func BenchmarkHoare(b *testing.B) {
	nums:=make([]int,len(nums1))
	copy(nums,nums1)
	for i := 0; i < b.N; i++ {
		QuickSort0(nums,0,len(nums)-1)
	}
}


func BenchmarkIOA0(b *testing.B) {
	nums:=make([]int,len(nums1))
	copy(nums,nums1)
	for i := 0; i < b.N; i++ {
		QuickSort1(nums,0,len(nums)-1)
	}
}

func BenchmarkIOA1(b *testing.B) {
	nums:=make([]int,len(nums1))
	copy(nums,nums1)
	for i := 0; i < b.N; i++ {
		QuickSort2(nums,0,len(nums)-1)
	}
}

func BenchmarkIOA2(b *testing.B) {
	nums:=make([]int,len(nums1))
	copy(nums,nums1)
	for i := 0; i < b.N; i++ {
		QuickSort3(nums,0,len(nums)-1)
	}
}

func BenchmarkRandNoSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(i+1)
	}
}
func BenchmarkRandWithSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(int64(i+1))
		rand.Intn(i+1)
	}
}

