package leetcode

import "math"

//贪心
func storeWater_0(bucket []int, vat []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	maxFillOps := 0
	for _, v := range vat {
		if v > maxFillOps {
			maxFillOps = v
		}
	}
	if maxFillOps == 0 {
		return 0
	}
	minOps := math.MaxInt64
	for fillOps := 1; fillOps <= maxFillOps && fillOps < minOps; fillOps++ {
		sumLvUpOps := 0
		for i := range bucket {
			// 所需升级次数 ceil(vat[i]/fillOps)-bucket[i]
			sumLvUpOps += max(0, (vat[i]+fillOps-1)/fillOps-bucket[i])
		}
		minOps = min(minOps, fillOps+sumLvUpOps)
	}
	return minOps
}

func storeWater_2(bucket []int, vat []int) int {
	ceil := func(x, y int) int {
		if x == 0 {
			return 0
		}
		q := x / y
		if x%y != 0 {
			q++
		}
		return q
	}

	minOps := math.MaxInt64
	//sumFillOps+currOps
	minSFOCO := math.MaxInt64
	lvUpOps := 0
	needLvUpIndex := -1
	for {
		maxFillOps := 0
		sumFillOps := 0
		if needLvUpIndex >= 0 {

			bucket[needLvUpIndex]++
			lvUpOps++
		}
		for i := range vat {
			if vat[i] != 0 && bucket[i] == 0 {
				bucket[i]++
				lvUpOps++
			}
			currFillOps := ceil(vat[i], bucket[i])
			sumFillOps += currFillOps
			if currFillOps > maxFillOps {
				maxFillOps = currFillOps
				needLvUpIndex = i
			}
		}
		currOps := lvUpOps + maxFillOps
		if currOps <= minOps {
			minOps = currOps
		}
		if minOps == 0 {
			break
		}
		if currSFOCO := sumFillOps + currOps; currSFOCO <= minSFOCO {
			minSFOCO = currSFOCO
		} else {
			break
		}
	}
	return minOps
}
