package leetcode


func IsPowerOfFour(n int) bool {
	mask:=0xaaaaaaaaaaaaaaa
	if n==0 || (n&mask==0 && n&(n-1)==0){
		return true
	}else{
		return false
	}
    

}
