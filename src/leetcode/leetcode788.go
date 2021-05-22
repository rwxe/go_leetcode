package leetcode

import (
	"fmt"
	"strconv"
)


func Spin(n int)int{
	if n==2{
		return 5
	}else if n==5{
		return 2
	}else if n==6{
		return 9
	}else if n==9{
		return 6
	}else if n==1{
		return 1
	}else if n==0{
		return 0
	}else if n==8{
		return 8
	}else{
		return -1
	}
}

func IsGoodDigits(n int)bool{
	nStr:=strconv.Itoa(n)
	buff:=""
	for _,char:=range nStr{
		rChar:=Spin(int(char-'0'))
		if rChar!=-1{
			buff+=fmt.Sprint(rChar)
		}else{
			return false
		}
		
	}
	//fmt.Printf("%T,%q\n",buff,buff)
	spinDig,_:=strconv.Atoi(buff)
	if spinDig!=n{
		return true
	}else{
		return false
	}
}

func RotatedDigits(N int) int {
	count:=0
	for i:=1;i<=N;i++{
		if IsGoodDigits(i){
			count+=1
		}
	}
	return count

}
