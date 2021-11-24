package leetcode

import "fmt"

type RangeFreqQuery struct {
	arr             []int
	rangeFreqMapMap map[int]map[int]int
	theMap          map[[2]int]int
}

type segmentTreeNode5186 struct {
	currentRangeMap map[int]int
}

func (rfq *RangeFreqQuery) pushup(lMap, rMap map[int]int) map[int]int {
	result := make(map[int]int)
	for k, v := range lMap {
		result[k] = v
	}
	for k, v := range rMap {
		result[k] += v
	}
	return result

}
func (rfq *RangeFreqQuery) dfs(l, r int) map[int]int {
	if l == r {
		return map[int]int{rfq.arr[l]: 1}
	}
	m := (l + r) / 2
	lMap := rfq.dfs(l, m)
	rMap := rfq.dfs(m+1, r)
	rfq.rangeFreqMapMap[r] = rfq.pushup(lMap, rMap)
	return rfq.rangeFreqMapMap[r]
}

func Constructor5186(arr []int) RangeFreqQuery {
	rfq := RangeFreqQuery{arr, make(map[int]map[int]int), make(map[[2]int]int)}

	rfq.dfs(0, len(rfq.arr)-1)
	for i, v := range rfq.rangeFreqMapMap {
		fmt.Println(i, v)
	}
	return rfq
}

func (rfq *RangeFreqQuery) Query(left int, right int, value int) int {
	if v, ok := rfq.theMap[[2]int{left, right}]; ok {
		return v
	}
	return 0
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
