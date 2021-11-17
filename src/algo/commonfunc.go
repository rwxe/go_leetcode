package algo

func MaxInts(items ...int) int {
	max := items[0]
	for _, v := range items {
		if v > max {
			max = v
		}
	}
	return max
}

func MinInts(items ...int) int {
	min := items[0]
	for _, v := range items {
		if v < min {
			min = v
		}
	}
	return min
}
