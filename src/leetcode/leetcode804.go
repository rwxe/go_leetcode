package leetcode

func uniqueMorseRepresentations(words []string) int {
	morseTable := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	theSet := make(map[string]struct{})
	for _, word := range words {
		str := ""
		for i := range word {
			str += morseTable[int(word[i]-'a')]
		}
		theSet[str] = struct{}{}
	}
	return len(theSet)
}
