package leetcode

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	iw, ic, jw, jc := 0, 0, 0, 0
	for iw < len(word1) && jw < len(word2) {

		if word1[iw][ic] != word2[jw][jc] {
			return false
		} else {
			ic++
			jc++
		}
		if ic >= len(word1[iw]) {
			iw++
			ic = 0
		}
		if jc >= len(word2[jw]) {
			jw++
			jc = 0
		}

	}
	if iw == len(word1) && jw == len(word2) {
		return true

	} else {
		return false
	}
}
