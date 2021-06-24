package leetcode


func bitLen(n int) int {
	count := 0
	if n < 1 {
		return count
	}
	for n != 0 {
		n >>= 1
		count++
	}
	return count
}

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	postive := true
	if n < 0 {
		n = -n
		postive = false
	}
	bitLen := bitLen(n)
	products := make([]float64, bitLen)
	products[0] = x
	for i := range products {
		if i == 0 {
			continue
		}
		products[i] = products[i-1] * products[i-1]
	}
	//fmt.Println(products)
	sum := 1.0
	for mask, i := 1, 0; n%mask != n; mask, i = mask<<1, i+1 {
		//		fmt.Println(sum)
		if mask&n != 0 {
			sum *= products[i]
		}

	}
	if !postive {
		return 1 / sum
	} else {
		return sum

	}

}
