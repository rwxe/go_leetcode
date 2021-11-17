package leetcode
func trailingZeroes(n int) int {
    count5:=0
	for n>0{
		count5+=n/5
		n/=5
	}
	return count5
}
func trailingZeroes_0(n int) int {
    count5:=0
    for i:=1;i<=n;i++{
        num:=i
        for num>0{
            if num%5==0{
                count5++
                num/=5
            }else{break}
        }
    }
    return count5
    
}
