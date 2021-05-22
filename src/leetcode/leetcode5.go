package leetcode

import "fmt"

// 动态规划
func LongestPalindrome(s string) string {
	begin:=0
	maxLen:=1
	if len(s)<2{
		return s
	}
	dp:=make([][]bool,len(s))
	for i:=0;i<len(s);i++{
		dp[i]=make([]bool,len(s))
	}
	// dp[i][j]==dp[i+1][j-1] && s[i]==s[j]
	// dp[i][j]==dp[i+1][j-1] && s[i]==s[j]
	for i:=0;i<len(s);i++{
		dp[i][i]=true
	}
	for j:=1;j<len(s);j++{
		for i:=0;i<j;i++{
			if s[i]!=s[j]{
				dp[i][j]=false
			}else{
				if j-i<3{
					dp[i][j]=true
				}else{
					dp[i][j]=dp[i+1][j-1]
				}
			}
			if dp[i][j]==true && j-i+1>maxLen{
				begin=i
				maxLen=j-i+1
			}
		}
	}
	fmt.Println(begin,maxLen)
	return s[begin:begin+maxLen]

}
