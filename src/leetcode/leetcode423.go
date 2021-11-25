package leetcode

import (
	"sort"
)

func OriginalDigits(s string) string {
	//0 1 2 3 4 5 6 7 8 9
	//zero one two three four five six seven eight nine
	//"z":0
	//"w":2
	//"u":4
	//"x":6
	//"g":8
	//"f":4,5
	//"h":3,8
	//"s":6,7
	//"v":5,7
	//"r":0,3,4
	//"t":2,3,8
	//"n":1,7,9
	//"o":0,1,2,4
	//"i":5,6,8,9
	//"e":0,1,3,5,7,8,9
	countMap := make(map[rune]int)
	for _, v := range s {
		countMap[v]++
	}
	count := [10]int{}
	count[0] = countMap['z']
	count[2] = countMap['w']
	count[4] = countMap['u']
	count[6] = countMap['x']
	count[8] = countMap['g']

	count[5] = countMap['f'] - count[4]
	count[3] = countMap['h'] - count[8]
	count[7] = countMap['s'] - count[6]

	count[1] = countMap['o'] - count[0] - count[2] - count[4]
	count[9] = countMap['i'] - count[5] - count[6] - count[8]
	result := make([]byte, 0)
	for i, v := range count {
		for j := 0; j < v; j++ {
			result = append(result, byte(i+'0'))
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return string(result)

}
