package leetcode

import (
	"math"
)

// 异或求对数
func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return left & right
	}
	exp := int(math.Log2(float64(left ^ right)))
	mask := 1<<(exp+1) - 1
	return right &^ mask

}

// lowbit操作
func rangeBitwiseAnd_0(left int, right int) int {
	for left < right {
		right &= right - 1
	}
	return right
}

// 异或求对数，快速求对数
func rangeBitwiseAnd_1(left int, right int) int {
	if left == right {
		return left & right
	}
	floatNum := float64(left ^ right)
	//exp := ((*(*int)(unsafe.Pointer(&floatNum)) >> 52) & 0x7ff)
	//1,11,52
	exp := (math.Float64bits(floatNum) >> 52) & 0x7ff
	if exp != 0 {
		exp -= 0x3ff
	}
	mask := 1<<(exp+1) - 1
	return right &^ mask
}
