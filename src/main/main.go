package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"src/leetcode"
	"unsafe"
)

func printAllocMain() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}
func main() {

	var words = []string{}
	for i := 'a'; i <= 'z'; i++ {
		// 将 rune 转换为字符串并添加到切片中
		words = append(words, string(i))
	}
	l := leetcode.Constructor359()
	for i := 0; i < 100; i++ {
		msg := words[rand.Intn(26)]
		l.ShouldPrintMessageLRU(i, msg)
	}
	runtime.GC()
	//单次测试的调试信息
	B := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int)(unsafe.Pointer(&l)))) + 9))
	fmt.Println("\nthe hmap B size", B)
	printAllocMain()

	//	for k := range l.LogHist {
	//		if delTimes, ok := l.DeletedKey[k]; ok {
	//			fmt.Println("WTF", k, delTimes)
	//		}
	//	}

}
