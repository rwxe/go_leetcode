package leetcode

func firstUniqChar_1(s string) byte {
	window := make([]byte, 0)
	showedUp := make(map[byte]int)
	for i := range s {
		if _, ok := showedUp[s[i]]; !ok {
			window = append(window, s[i])
			showedUp[s[i]] = i
		} else {
			showedUp[s[i]] = -1
			for len(window) != 0 && showedUp[window[0]] == -1 {
				window = window[1:]
			}
		}
	}
	if len(window) != 0 {
		return window[0]
	} else {
		return ' '
	}
}
func firstUniqChar_0(s string) byte {
	showedUp := make(map[byte]int)
	for i := range s {
		showedUp[s[i]]++
	}
	for i := range s {
		if showedUp[s[i]] < 2 {
			return s[i]
		}
	}
	return ' '

}
