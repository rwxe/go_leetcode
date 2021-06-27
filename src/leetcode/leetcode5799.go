package leetcode
//TODO

func isWonderful(subStr string)bool{
	freq:=make(map[rune]int)
	count:=0
	for _,v:=range subStr{
		freq[v]++
	}
	for _,v:=range freq{

		if v%2!=0{
			count++
		}
	}
	if count>1{
		return false
	}else{
		return true
	}


}

func WonderfulSubstrings(word string) int64 {
	count:=0
	for i:=0;i<len(word);i++{
		for j:=0;j<len(word)-i;j++{
			subStr:=word[j:j+i+1]
			if isWonderful(subStr){
				count++

			}
		}
	}

	return int64(count)
}
