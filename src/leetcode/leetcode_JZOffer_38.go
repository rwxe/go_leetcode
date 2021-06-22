package leetcode

import (
	"sort"
)

type runeSlice []rune
func (rs runeSlice)Len()int{ return len(rs) }
func (rs runeSlice)Swap(i,j int){rs[i],rs[j]=rs[j],rs[i]}
func (rs runeSlice)Less(i,j int)bool{return rs[i]<rs[j]}

func inJZO38(list []rune,elem rune)bool{
	for _,v:=range list{
		if v==elem{
			return true
		}
	}
	return false
}


func btJZO38(nums ,path []rune , length int,result *[]string) {
	if len(path) >= length {
		*result = append(*result, string(append([]rune(nil), path...)))
	}
	for i := 0; i < len(nums); i++ {
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		path = append(path, nums[i])
		newNums:=append([]rune(nil),nums[:i]...)
		newNums=append(newNums,nums[i+1:]...)
		btJZO38(newNums, path, length,result)
		path = path[:len(path)-1]
	}

}

func Permutation(s string) []string {
	result := make([]string, 0)
	sSlice:=runeSlice(s)
	sort.Sort(sSlice)
	btJZO38(sSlice,[]rune{},len(s),&result)
	return result

}
