package leetcode

func ReverseVowels(s string) string {
	byteArr := []byte(s)
	stack := make([]byte, 0)
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {}}
	for i := 0; i < len(byteArr); i++ {
		if _, ok := vowels[byteArr[i]]; ok {
			stack = append(stack, byteArr[i])
		}
	}
	for i := len(byteArr) - 1; i >= 0; i-- {
		if _, ok := vowels[byteArr[i]]; ok {
			byteArr[i] = stack[0]
			stack = stack[1:]
		}
	}
	return string(byteArr)

}
func ReverseVowels_1(s string) string {
	byteArr := []byte(s)
	i, j := 0, len(byteArr)-1
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {}}
	in := func(char byte) bool {
		_, ok := vowels[char]
		return ok
	}
	for i < j {
		for ; i < len(byteArr) && !in(byteArr[i]); i++ {
		}
		for ; j >= 0 && !in(byteArr[j]); j-- {
		}
		if i < j {
			byteArr[i], byteArr[j] = byteArr[j], byteArr[i]
			i++
			j--
		}
	}
	return string(byteArr)
}
