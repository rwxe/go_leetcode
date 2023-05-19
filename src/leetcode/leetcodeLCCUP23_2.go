package leetcode

import (
	"strings"
)

func AdventureCamp(expeditions []string) int {
	vistedCamps := make(map[string]struct{}, len(expeditions[0]))
	maxCount := 0
	minI := -1
	var count int
	for i, e := range expeditions {
		camps := strings.Split(e, "->")
		if i == 0 {
			for _, c := range camps {
				vistedCamps[c] = struct{}{}
			}
		} else {
			campsSet := make(map[string]struct{}, len(camps))
			for _, c := range camps {
				campsSet[c] = struct{}{}
			}
			count = 0
			for c := range campsSet {
				var ok bool
				if _, ok = vistedCamps[c]; !ok && c != "" {
					count++
				}
				vistedCamps[c] = struct{}{}
				//fmt.Println(c, ok, count)
			}
			//fmt.Println(initCamps, campsSet, count, maxCount)

		}
		if count > maxCount {
			maxCount = count
			minI = i
		}
	}
	return minI
}
