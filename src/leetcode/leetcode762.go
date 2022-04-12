package leetcode

func countPrimeSetBits_1(left int, right int) int {

	//2, 3, 5, 7, 11, 13, 17, 19
	count1 := func(num int) int {
		count := 0
		for ; num != 0; num &= num - 1 {
			count++
		}
		return count
	}
	primesMask := 2693408940
	result := 0
	for i := left; i <= right; i++ {
		if (1<<count1(i))&primesMask != 0 {
			result++
		}
	}
	return result
}
func countPrimeSetBits_0(left int, right int) int {
	//2, 3, 5, 7, 11, 13, 17, 19
	count1 := func(num int) int {
		count := 0
		for ; num != 0; num &= num - 1 {
			count++
		}
		return count
	}

	primes := map[int]bool{
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
	}
	result := 0
	for i := left; i <= right; i++ {
		if primes[count1(i)] {
			result++
		}
	}
	return result
}
