package leetcode

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var result int64
	max1 := total / cost1
	for i := 0; i <= max1; i++ {
		result += int64((total-i*cost1)/cost2 + 1)
	}
	return result
}
