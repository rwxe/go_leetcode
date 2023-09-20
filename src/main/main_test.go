package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"src/leetcode"
	"testing"
	"unsafe"
)

var words = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
	"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
	"31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
	"41", "42", "43", "44", "45", "46", "47", "48", "49", "50",
}

// 内存占用打印
func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func BenchmarkDefault(b *testing.B) {
	l := leetcode.Constructor359()
	for i := 0; i < 100_000; i++ {
		msg := words[rand.Intn(50)] + words[rand.Intn(50)] + words[rand.Intn(50)]
		l.ShouldPrintMessageDefault(i, msg)
	}
	runtime.GC()
	//单次测试的调试信息
	B := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int)(unsafe.Pointer(&l)))) + 9))
	fmt.Println("\nthe hmap B size", B)
	printAlloc()
	l.PrintDebugMaxLen()
}

func BenchmarkOptimized(b *testing.B) {
	b.ResetTimer()
	l := leetcode.Constructor359()
	for i := 0; i < 100_000; i++ {
		msg := words[rand.Intn(50)] + words[rand.Intn(50)] + words[rand.Intn(50)]
		l.ShouldPrintMessageLRU(i, msg)
	}
	runtime.GC()
	//单次测试的调试信息
	B := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int)(unsafe.Pointer(&l)))) + 9))
	fmt.Println("\nthe hmap B size", B)
	printAlloc()
	l.PrintDebugMaxLen()
}
