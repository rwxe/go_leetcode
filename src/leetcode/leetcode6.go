package leetcode

func convert6_2(s string, numRows int) string {
	// 0             0+t                    0+2t                     0+3t
	// 1      t-1    1+t            0+2t-1  1+2t            0+3t-1   1+3t
	// 2  t-2        2+t  0+2t-2            2+2t  0+3t-2             2+3t
	// 3             3+t                    3+2t                     3+3t
	if numRows == 1 {
		return s
	}
	result := make([]byte, 0, len(s))
	t := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; i+j < len(s); j += t {
			result = append(result, s[i+j])
			if 0 < i && i < numRows-1 && j+t-i < len(s) {
				result = append(result, s[j+t-i])
			}
		}
	}
	return string(result)
}
func convert6(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := make([]byte, 0, len(s))
	rows := make([][]byte, numRows)
	z := 0
	incr := false
	for i := range s {
		rows[z] = append(rows[z], s[i])
		if z == 0 || z == numRows-1 {
			incr = !incr
		}
		if incr {
			z++
		} else {
			z--
		}
	}
	for i := range rows {
		result = append(result, rows[i]...)
	}
	return string(result)
}
