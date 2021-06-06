package leetcode

func Subset474(strs []string)[][]string{
	result:=make([][]string,1)
	for i:=0;i<len(strs);i++{
		resultLen:=len(result)
		for j:=0;j<resultLen;j++{
			list:=append([]string(nil),result[j]...)
			list =append(list,strs[i])
			result=append(result,list)
		}
	}
	return result
}
func FindMaxForm(strs []string, m int, n int) int {
	//TODO 
	return 0


}
