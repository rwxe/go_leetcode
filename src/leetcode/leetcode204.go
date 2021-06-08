package leetcode

//import "math"

func IsPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%6 != 1 && n%6 != 5 {
		return false
	}
//  	sqrt := int(math.Sqrt(float64(n)))
//  	for i := 5; i <= sqrt; i += 6 {
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func CountPrimes0(n int) int {
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

func CountPrimes1(n int) int {
	count := 0
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
			for j := 2 * i; j < n; j+=i{
				isPrime[j] = false
			}

		}
	}
	return count
}
