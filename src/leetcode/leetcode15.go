package leetcode

import (
	_"fmt"
	"reflect"
	"sort"
)

func inResult15(result [][]int, s []int) bool {
	for _,set:=range result{
		if reflect.DeepEqual(set,s){
			return true
		}else{
			//fmt.Println(set,"一样",s)
		}
	}
	return false

}

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)

	for i, n1 := range nums {
		map1 := make(map[int]int)
		traget := 0 - n1
		for j, n2 := range nums {
			if j == i {
				continue
			}
			if k, ok := map1[traget-n2]; ok {
				set := []int{n1, n2, nums[k]}
				//fmt.Println("排序前",set)
				sort.Ints(set)
				//fmt.Println("排序后",set)
				if !inResult15(result, set) {
					result = append(result, set)
				}
			} else {
				map1[n2] = j
			}
		}

	}
	return result

}
