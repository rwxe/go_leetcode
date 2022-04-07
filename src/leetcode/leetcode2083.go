package leetcode

func numberOfSubstrings_2(s string) int64 {
	themap := [26]int{}
	result := 0
	for i := range s {
		themap[s[i]-'a']++
		result += themap[s[i]-'a']
	}
	return int64(result)
}
func numberOfSubstrings_1(s string) int64 {
	theMap := make(map[byte]int, 26)
	result := 0
	for i := range s {
		theMap[s[i]]++
		result += theMap[s[i]]
	}
	return int64(result)
}
func numberOfSubstrings_0(s string) int64 {
	var count int64 = int64(len(s))
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				count++
			}
		}
	}
	return count
}
