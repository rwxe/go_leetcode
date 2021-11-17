package leetcode

import (
	"sort"
	"strconv"
)

func GetHint_0(secret string, guess string) string {
	bulls, cows := 0, 0
	secretMap := make(map[int]int)
	guessMap := make(map[int]int)
	min := func(i, j int) int {
		if i < j {
			return i
		} else {
			return j
		}
	}
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretMap[int(secret[i]-'0')]++
			guessMap[int(guess[i]-'0')]++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(secretMap[i], guessMap[i])
	}
	result := strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
	return result
}
func GetHint_1(secret string, guess string) string {
	bulls, cows := 0, 0
	secretBytes := []byte(secret)
	guessBytes := []byte(guess)
	for i, j := len(secretBytes)-1, len(guessBytes)-1; i >= 0 && j >= 0; {
		if secretBytes[i] == guessBytes[j] {
			bulls++
			secretBytes[i] = 0
			guessBytes[j] = 0
		}
		i, j = i-1, j-1
	}
	lessFunc1 := func(i, j int) bool { return secretBytes[i] < secretBytes[j] }
	lessFunc2 := func(i, j int) bool { return guessBytes[i] < guessBytes[j] }
	sort.Slice(secretBytes, lessFunc1)
	sort.Slice(guessBytes, lessFunc2)
	for i, j := 0, 0; i < len(secretBytes) && j < len(guessBytes); {
		if secretBytes[i] == byte(0) {
			i++
		} else if guessBytes[j] == byte(0) {
			j++
		} else if secretBytes[i] == guessBytes[j] {
			cows++
			i++
			j++
		} else if secretBytes[i] > guessBytes[j] {
			j++
		} else {
			i++
		}
	}
	result := strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
	return result

}
