package leetcode

func hasAlternatingBits_2(n int) bool {
	return (n^n>>1)&((n^n>>1)+1) == 0
}
func hasAlternatingBits(n int) bool {
	p := 1
	flag := true
	if n&p == 1 {
		flag = true
	} else {
		flag = false
	}
	for n != 0 {
		if n&p == 1 && !flag {
			return false
		} else if n&p == 0 && flag {
			return false
		}
		n >>= 1
		flag = !flag
	}
	return true
}
