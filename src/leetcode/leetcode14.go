package leetcode
func LongestCommonPrefix(strs []string) string {
	if strs==nil{
		return ""
	}
	prefix:=""
	minLen:=len(strs[0])
	minStr:=strs[0]
	for _,str:=range strs{
		if len(str)<minLen{
			minStr=str
		}
	}
	for i:=range minStr{
		for _,str:=range strs{
			if str[i]!=minStr[i]{
				return prefix
			}
		}
		prefix+=string(minStr[i])
	}
	return prefix


}
