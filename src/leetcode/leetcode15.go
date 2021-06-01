package leetcode

import (
	"fmt"
	_ "fmt"
	"reflect"
	"sort"
)

func inResult15(result [][]int, s []int) bool {
	for _, set := range result {
		if reflect.DeepEqual(set, s) {
			return true
		} else {
			//fmt.Println(set,"一样",s)
		}
	}
	return false

}

func ThreeSum0(nums []int) [][]int {
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

func ThreeSum1(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			k := len(nums) - 1
			for j := i + 1; j < len(nums)-1; j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					for nums[i]+nums[j]+nums[k] > 0  && j<k{
						k -= 1
					}
					if j==k{
						break
					}
					if nums[i]+nums[j]+nums[k] == 0 {
						result = append(result, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}

	}
	return result

}
