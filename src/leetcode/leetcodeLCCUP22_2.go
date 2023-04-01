package leetcode

func TransportationHub(path [][]int) int {
	type XYSum struct {
		X, Y int
	}
	nodeXYSum := make(map[int]*XYSum)
	nodeCount := make(map[int]struct{})

	for _, p := range path {
		from, to := p[0], p[1]
		nodeCount[from] = struct{}{}
		nodeCount[to] = struct{}{}

		fromXYSum, ok := nodeXYSum[from]
		if !ok {
			nodeXYSum[from] = &XYSum{}
			fromXYSum = nodeXYSum[from]
		}
		toXYSum, ok := nodeXYSum[to]
		if !ok {
			nodeXYSum[to] = &XYSum{}
			toXYSum = nodeXYSum[to]
		}
		fromXYSum.X++
		toXYSum.Y++
	}
	for k, n := range nodeXYSum {
		if n.X == 0 && n.Y == len(nodeCount)-1 {
			return k
		}
	}
	return -1
}
