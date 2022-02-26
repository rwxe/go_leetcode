package leetcode

func CountEven(num int) int {
	sumIsEven := func(num int) bool {
		tempSum := 0
		for num != 0 {
			tempSum += num % 10
			num /= 10
		}
		if tempSum%2 == 0 {
			return true
		} else {
			return false
		}
	}
	count := 0
	for i := 1; i <= num; i++ {
		if sumIsEven(i) {
			count++
		}
	}
	return count
}
