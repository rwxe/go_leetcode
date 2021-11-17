package leetcode

import "src/algo"

func minDistance(word1 string, word2 string) int {
	return dp72_2(word1, word2)
}

func dp72_2(word1, word2 string) int {
	//dp[i,j]代表word1[0:i],word2[0:j]的最短编辑距离
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1)
	for i := range dp {
		dp[i] = make([]int, n2)
	}
	for i := 0; i < n1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n2; j++ {
		dp[0][j] = j
	}
	for i := 1; i < n1; i++ {
		for j := 1; j < n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = algo.MinInts(
					//add
					dp[i][j-1],
					//del
					dp[i-1][j],
					//add
					dp[i-1][j-1],
				) + 1
			}
		}
	}
	return dp[n1-1][n2-1]

}

func dp72_1(memory map[[2]int]int, word1, word2 string, i, j int) int {
	//dp(word1,word2,i,j)代表word1[0:i+1],word2[0:j+1]的最短编辑距离
	if i == -1 {
		return j + 1
	} else if j == -1 {
		return i + 1
	} else if v, ok := memory[[2]int{i, j}]; ok {
		return v
	}

	if word1[i] == word2[j] {
		memory[[2]int{i, j}] = dp72_1(memory, word1, word2, i-1, j-1)
	} else {
		memory[[2]int{i, j}] = algo.MinInts(
			//add
			dp72_1(memory, word1, word2, i, j-1),
			//del
			dp72_1(memory, word1, word2, i-1, j),
			//replace
			dp72_1(memory, word1, word2, i-1, j-1),
		) + 1
	}
	return memory[[2]int{i, j}]

}
func dp72_0(word1, word2 string, i, j int) int {
	if i == -1 {
		return j + 1
	} else if j == -1 {
		return i + 1
	} else if word1[i] == word2[j] {
		return dp72_0(word1, word2, i-1, j-1)
	} else {
		return algo.MinInts(
			//add
			dp72_0(word1, word2, i, j-1),
			//del
			dp72_0(word1, word2, i-1, j),
			//replace
			dp72_0(word1, word2, i-1, j-1),
		) + 1
	}
}
