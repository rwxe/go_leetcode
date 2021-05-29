package leetcode
// 计算非负数的二进制1数量
func BrianKernighanBitCount(num int)int{
	count:=0
	for num>0{
		num&=num-1
		count+=1
	}
	return count

}

func CountBits(n int) []int {
	dp:=make([]int,n+1)
	highBit:=0
	for i:=1;i<=n;i++{
		if i&(i-1)==0{
			highBit=i
			dp[i]=1
		}else{
			dp[i]=dp[i-highBit]+1
		}

	}
	return dp

}
