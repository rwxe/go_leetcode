package leetcode

func RemoveKdigits(num string, k int) string {
	srcLen:=len(num)
	dstLen:=srcLen-k
	for time:=0;time<k;time++{
		for i:=1;i<len(num);i++{
			if num[i-1]>num[i]{
				num=num[:i-1]+num[i:]
				break
				//记得这里要break
			}
		}
	}
	if len(num)>dstLen{
		num=num[:dstLen]
	}
	if len(num)==0{
		return "0"
	}
	for num[0]=='0' && len(num)!=1{
		num=num[1:]
	}
	return num

}
