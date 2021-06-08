package leetcode

import "math"


func IsPrime(n int) bool {
	if n<=3{
		return n>1
	}
	if n%6 != 1 && n%6 != 5 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrt; i+=6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	count := 0
	if n < 1 {
		return 0
	}
	for i := 0; i < n; i++ {
		if IsPrime(i) {
			count++
		}
	}
	return count

}
