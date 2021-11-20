package leetcode

import "src/algo"

func minCost(costs [][]int) int {
	//回溯
	//result:=1<<31
	//bt(costs,0,0,-1,&result)
	//return result

	// 记忆化搜索
	//memory:=make([][]int,len(costs))
	//for i:=range memory{
	//    memory[i]=make([]int,len(costs[i]))
	//    for j:=range memory[i]{
	//        memory[i][j]=-1
	//    }
	//}
	//return dfs(costs,-1, -1, memory)

	// 动态规划
	return dp(costs)
}

func dp(costs [][]int) int {
	tempCosts := make([]int, 0, len(costs[0]))
	for i := len(costs) - 2; i >= 0; i-- {
		for j := 0; j < len(costs[i]); j++ {
			tempCosts := tempCosts[:0]
			for k := 0; k < len(costs[i]); k++ {
				if k == j {
					continue
				}
				tempCosts = append(tempCosts, costs[i+1][k])
			}
			costs[i][j] += algo.MinInts(tempCosts...)

		}
	}
	return algo.MinInts(costs[0]...)
}

func dfs256(costs [][]int, houseIndex, colorIndex int, memory [][]int) int {
	if houseIndex == len(costs) {
		return 0
	}
	currentCost := 0
	if houseIndex >= 0 {
		currentCost = costs[houseIndex][colorIndex]
		if v := memory[houseIndex][colorIndex]; v != -1 {
			return v
		}
	}

	tempCosts := make([]int, 0, len(costs[0]))
	for i := 0; i < len(costs[0]); i++ {
		if i == colorIndex {
			continue
		}
		tempCosts = append(tempCosts, dfs256(costs, houseIndex+1, i, memory))
	}
	currentCost += algo.MinInts(tempCosts...)
	if houseIndex >= 0 {
		memory[houseIndex][colorIndex] = currentCost

	}
	return currentCost
}

func bt256(costs [][]int, houseIndex int, currentCost, preColor int, result *int) {
	if houseIndex == len(costs) {
		if *result > currentCost {
			*result = currentCost
		}
		return
	} else if currentCost >= *result {
		return
	}
	for i := 0; i < len(costs[houseIndex]); i++ {
		if i == preColor {
			continue
		}
		currentCost += costs[houseIndex][i]
		bt256(costs, houseIndex+1, currentCost, i, result)
		currentCost -= costs[houseIndex][i]
	}
}
