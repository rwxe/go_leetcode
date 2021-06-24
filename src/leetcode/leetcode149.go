package leetcode

func k149(p1, p2 []int) float64 {
	return float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
}

func MaxPoints(points [][]int) int {
		maxLine := 0
	for i, p1 := range points {
		pointsWithoutP1 := append([][]int(nil), points[:i]...)
		pointsWithoutP1 = append(pointsWithoutP1, points[i+1:]...)
		kMap := make(map[float64]int)
		for _, p2 := range pointsWithoutP1 {
			kMap[k149(p1, p2)]++
			if kMap[k149(p1, p2)] > maxLine {
				maxLine = kMap[k149(p1, p2)]
			}

		}

	}
	return maxLine+1

}
