package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func findRestaurant_1(list1 []string, list2 []string) []string {
	set1 := make(map[string]int, len(list1))
	minIndexSum := math.MaxInt64
	for i1, v := range list1 {
		set1[v] = i1
	}
	result := make([]string, 0)
	for i2, v2 := range list2 {
		if i1, ok := set1[v2]; ok {
			if i1+i2 < minIndexSum {
				minIndexSum = i1 + i2
				result = []string{v2}
			} else if i1+i2 == minIndexSum {
				result = append(result, v2)
			} else if i2 > minIndexSum {
				//优化
				break

			}
		}
	}
	return result
}
func findRestaurant_0(list1 []string, list2 []string) []string {
	set1 := make(map[string]int)
	//[交集字符串在list1的下标，下标和]
	intersection := make([][2]int, 0)
	for i1, v := range list1 {
		set1[v] = i1
	}
	for i2, v := range list2 {
		if i1, ok := set1[v]; ok {
			intersection = append(intersection, [2]int{i1, i1 + i2})
		}
	}
	sort.Slice(intersection, func(i, j int) bool {
		return intersection[i][1] < intersection[j][1]
	})
	fmt.Println(intersection)
	minIndexSum := intersection[0][1]
	indexLimit := 0
	for i, v := range intersection {
		if v[1] <= minIndexSum {
			indexLimit = i
		} else {
			break
		}
	}
	result := make([]string, 0)
	for i := 0; i <= indexLimit; i++ {
		result = append(result, list1[intersection[i][0]])

	}
	return result

}
