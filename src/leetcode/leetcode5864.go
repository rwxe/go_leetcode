package leetcode

import "sort"

func NumberOfWeakCharacters(properties [][]int) int {
	powered := make([][]int, len(properties))
	count:=0
	for i, v := range properties {
		powered[i] = []int{v[0]*10 + v[1], v[0],v[1]}
	}
	sort.Slice(powered, func(i, j int) bool { return powered[i][0] < powered[j][0] })
	for i := 0; i < len(powered); i++ {
		for j:=i+1;j<len(powered);j++{
			if powered[j][1]>powered[i][1] && powered[j][2]>powered[i][2]{
				count++
				break
			}
		}

	}
	return count

}
