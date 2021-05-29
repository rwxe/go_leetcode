package leetcode

func lengthOfLongestSubstring(s string) int {
	window := make([]rune, 0)
	existed:=make(map[rune]bool)
	maxLen:=0
	for _,char:=range s{
		for existed[char]{
			delete(existed,window[0])
			window=window[1:]
		}
		window=append(window,char)
		existed[char]=true
		if maxLen<len(window){
			maxLen=len(window)
		}
	}
	return maxLen

}
