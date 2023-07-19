package leetcode

import (
	"math"
	"sync"
)

var result1828 []int
var wg sync.WaitGroup

func distance(p1x, p1y, p2x, p2y int) float64 {
	return math.Sqrt(float64(math.Pow(float64(p1x-p2x), 2) + math.Pow(float64(p1y-p2y), 2)))
}
func gogogo(p1x, p1y, p2x, p2y, r, i int) {
	if distance(p1x, p1y, p2x, p2y) <= float64(r) {
		result1828[i] += 1
	}
	wg.Done()

}
func CountPoints(points [][]int, queries [][]int) []int {
	result1828 = make([]int, len(queries))
	for i, circle := range queries {
		for _, p := range points {
			//			if distance(p[0], p[1], circle[0], circle[1]) <= float64(circle[2]) {
			//				result[i] += 1
			//
			//			}
			wg.Add(1)
			go gogogo(p[0], p[1], circle[0], circle[1], circle[2], i)
		}
	}

	wg.Wait()
	return result1828

}
func CountPoints2(points [][]int, queries [][]int) []int {
	result := make([]int, len(queries))
	for i, circle := range queries {
		for _, p := range points {
			if distance(p[0], p[1], circle[0], circle[1]) <= float64(circle[2]) {
				result[i] += 1

			}
		}
	}

	return result

}
