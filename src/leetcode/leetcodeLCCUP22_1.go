package leetcode

func temperatureTrend(temperatureA []int, temperatureB []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	const (
		DOWN   = -1
		STABLE = 0
		UP     = +1
	)
	trendA := make([]int, len(temperatureA)-1)
	trendB := make([]int, len(temperatureB)-1)
	for i := range trendA {
		if temperatureA[i+1] < temperatureA[i] {
			trendA[i] = DOWN
		} else if temperatureA[i+1] == temperatureA[i] {
			trendA[i] = STABLE
		} else {
			trendA[i] = UP
		}
	}
	for i := range trendB {
		if temperatureB[i+1] < temperatureB[i] {
			trendB[i] = DOWN
		} else if temperatureB[i+1] == temperatureB[i] {
			trendB[i] = STABLE
		} else {
			trendB[i] = UP
		}
	}
	sameDays := 0
	maxSamDays := 0
	minLen := min(len(trendA), len(trendB))
	for i := 0; i < minLen; i++ {
		if trendA[i] == trendB[i] {
			sameDays++
		} else {
			sameDays = 0
		}
		maxSamDays = max(maxSamDays, sameDays)

	}
	return maxSamDays

}
