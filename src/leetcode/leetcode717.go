package leetcode
func IsOneBitCharacter(bits []int) bool {
	i:=0
	for i<len(bits)-1{
		if bits[i]==1{
			i+=2
		}else{
			i+=1
		}
	}
	if i==len(bits)-1{
		return true
	}else{
		return false
	}

}
