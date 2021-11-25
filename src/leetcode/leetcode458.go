package leetcode

import (
	"fmt"
	"math"
)

func PoorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// 香农信息熵公式
	// H(X)=-Sum(P(xi)log2(P(xi)))

	// Hb:=1000*(1/1000)math.Log2(1/1000)
	Hb := -math.Log2((1.0 / float64(buckets)))
	PMFi := float64(minutesToTest/minutesToDie) + 1.0
	PMFi = 1 / PMFi
	temp := -math.Log2(PMFi)
	fmt.Println(Hb, PMFi, temp)
	return int(math.Ceil(Hb / temp))

}
