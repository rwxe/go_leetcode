package leetcode

func maxProfit(prices []int) int {
	dpPre, dpCurrent := 0, 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
		if v := prices[i] - minPrice; v > dpPre {
			dpCurrent = v
			dpPre = dpCurrent
		} else {
			dpCurrent = dpPre
		}
	}
	return dpCurrent

}
