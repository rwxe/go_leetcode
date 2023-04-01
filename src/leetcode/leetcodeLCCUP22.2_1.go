package leetcode

func MinNumBooths(demand []string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	set := make(map[byte]int)
	sum := 0
	for _, v := range demand {
		tempSet := make(map[byte]int)
		for i := range v {
			tempSet[v[i]]++
		}
		for k, v := range tempSet {
			set[k] = max(set[k], v)
		}
	}
	for _, v := range set {
		sum += v
	}
	//log.Println(set)
	return sum
}
