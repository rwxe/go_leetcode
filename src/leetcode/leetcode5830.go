package leetcode

import "fmt"

func IsThree(n int) bool {
	count:=0
	for i:=1;i<=n;i++{
		if n%i==0{
			count++
		}
	}
	fmt.Println(count)
	return count==3

}
