package leetcode

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
