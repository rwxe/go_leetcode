package leetcode

func firstUniqChar_1(s string) byte {
	queue := make([]byte, 0)
	showedUp := make(map[byte]int)
	for i := range s {
		if _, ok := showedUp[s[i]]; !ok {
			queue = append(queue, s[i])
			showedUp[s[i]] = i
		} else {
			showedUp[s[i]] = -1
			for len(queue) != 0 && showedUp[queue[0]] == -1 {
				queue = queue[1:]
			}
		}
	}
	if len(queue) != 0 {
		return queue[0]
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
