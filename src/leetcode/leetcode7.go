package leetcode

import (
	"strconv"
)


func Reverse(x int) int{
	isNegative:=false
	if x<0{
		isNegative=true
		x=-x
	}
	xstr:=strconv.Itoa(x)
	strAry:=[]rune(xstr)
	for i,j:=0,len(xstr)-1;i<(len(xstr)/2);i,j=i+1,j-1{
		strAry[i],strAry[j]=strAry[j],strAry[i]
	}
	xstr=string(strAry)
	//fmt.Printf(xstr)
	x,_=strconv.Atoi(xstr)
	if isNegative{
		x=-x
	}
	if x>(1<<31) || x<(-1<<31){
		return 0
	}
	return x
	
}
