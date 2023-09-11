package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"src/leetcode"
	"testing"
)

var words = []string{
	"apple", "banana", "cherry", "date", "elderberry",
	"fig", "grape", "honeydew", "kiwi", "lemon",
	"mango", "orange", "peach", "quince", "raspberry",
	"strawberry", "tangerine", "uva", "watermelon", "xylophone",
	"yam", "zucchini", "avocado", "blueberry", "carrot",
	"dragonfruit", "eggplant", "flan", "grapefruit", "huckleberry",
	"icecream", "jackfruit", "kiwifruit", "leek", "melon",
	"nectarine", "olive", "pineapple", "quail", "radish",
	"starfruit", "tomato", "ugli", "vanilla", "walnut",
	"ximenia", "yizdu", "ziti", "golang", "python",
}

//内存占用打印
func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func BenchmarkDefault(b *testing.B) {
	b.ResetTimer()
	l := leetcode.Constructor359()
	for i := 0; i < 10000; i++ {
		msg := words[rand.Intn(50)] + words[rand.Intn(50)] + words[rand.Intn(50)]
		l.ShouldPrintMessageDefault(i, msg)
	}
	//单次测试的调试信息
	//B := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int)(unsafe.Pointer(&l)))) + 9))
	//fmt.Println("\nthe hmap B size", B)
	//printAlloc()
	//l.PrintMaxLen()
}

func BenchmarkOptimized(b *testing.B) {
	b.ResetTimer()
	l := leetcode.Constructor359()
	for i := 0; i < 10000; i++ {
		msg := words[rand.Intn(50)] + words[rand.Intn(50)] + words[rand.Intn(50)]
		l.ShouldPrintMessageLRU(i, msg)
	}
	//单次测试的调试信息
	//B := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int)(unsafe.Pointer(&l)))) + 9))
	//fmt.Println("\nthe hmap B size", B)
	//printAlloc()
	//l.PrintMaxLen()
}
