package leetcode

func isAlienSorted_2(words []string, order string) bool {
	ordermap := make([]int, 26)
	for i, c := range order {
		ordermap[c-'a'] = i
	}
	less := func(wordi, wordj string) bool {
		for i, j := 0, 0; i < len(wordi) || j < len(wordj); i, j = i+1, j+1 {
			if i == len(wordi) {
				return true
			} else if j == len(wordj) {
				return false
			}
			ci, cj := wordi[i]-'a', wordj[j]-'a'
			if ordermap[ci] < ordermap[cj] {
				return true
			} else if ordermap[ci] > ordermap[cj] {
				return false
			}
		}
		return true
	}

	for i := 1; i < len(words); i++ {
		if !less(words[i-1], words[i]) {
			return false
		}
	}
	return true
}
func isAlienSorted(words []string, order string) bool {
	ordermap := make(map[byte]int)
	for i := range order {
		ordermap[order[i]] = i
	}
	less := func(wordi, wordj string) bool {
		for i, j := 0, 0; i < len(wordi) || j < len(wordj); i, j = i+1, j+1 {
			if i == len(wordi) {
				return true
			} else if j == len(wordj) {
				return false
			}
			ci, cj := wordi[i], wordj[j]
			if ordermap[ci] < ordermap[cj] {
				return true
			} else if ordermap[ci] > ordermap[cj] {
				return false
			}
		}
		return true
	}
	for i := 1; i < len(words); i++ {
		if !less(words[i-1], words[i]) {
			return false
		}
	}
	return true
}
