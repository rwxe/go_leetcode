package leetcode

func SumOfMultiples(n int) int {
	//s=na1+n(n-1)d/2
	//s=n(n+1)d/2
	sumAS := func(m int) int {
		k := n / m
		return (k * (k + 1) * m) / 2
	}
	sum := sumAS(3) + sumAS(5) + sumAS(7) -
		sumAS(3*5) - sumAS(3*7) - sumAS(5*7) +
		sumAS(3*5*7)

	return sum
}
