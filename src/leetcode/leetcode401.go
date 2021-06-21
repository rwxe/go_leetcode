package leetcode

import (
	"strconv"
)


func bk401(num int)int{
	count:=0
	for num!=0{
		num&=num-1
		count++
	}
	return count

}

func numToTime(h,m int)string{
	hStr:=strconv.Itoa(h)
	mStr:=strconv.Itoa(m)
	if m<10{
		mStr="0"+mStr
	}
	return hStr+":"+mStr
}


func ReadBinaryWatch(turnedOn int) []string {
	// 0000:000000
	result := make([]string, 0)
	for h:=0;h<12;h++{
		for m:=0;m<60;m++{
			if bk401(h)+bk401(m)==turnedOn{
				result=append(result,numToTime(h,m))
			}
		}
	}
	return result

	

}
