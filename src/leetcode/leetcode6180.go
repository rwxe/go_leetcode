package leetcode

func smallestEvenMultiple(n int) int {
	var big, small int
	if n > 2 {
		big = n
		small = 2
	} else {
		big = 2
		small = n
	}
	if big%small == 0 {
		return big
	} else {
		return big * 2
	}
}
