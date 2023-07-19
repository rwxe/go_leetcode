package leetcode

import (
	"math"
)

func robotSim_1(commands []int, obstacles [][]int) int {
	//X:Y1,Y2...
	obMap := make(map[int]struct{}, len(obstacles))
	const POS_RANGE int = 60001
	for _, pos := range obstacles {
		obMap[pos[0]*POS_RANGE+pos[1]] = struct{}{}
	}
	direction := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	currFace := 0
	currPos := []int{0, 0}
	tempPos := []int{0, 0}

	move := func(steps int) {
		for i := 0; i < steps; i++ {
			copy(tempPos, currPos)
			tempPos[0] += direction[currFace][0]
			tempPos[1] += direction[currFace][1]
			if _, ok := obMap[tempPos[0]*POS_RANGE+tempPos[1]]; ok {
				break
			}
			copy(currPos, tempPos)
		}
	}
	turn := func(cmd int) {
		if cmd == -1 {
			currFace = (currFace + 1) % 4
		} else if cmd == -2 {
			currFace = (currFace + 3) % 4
		}
	}
	maxDistance := math.MinInt64
	for _, cmd := range commands {
		if cmd < 0 {
			turn(cmd)
		} else {
			move(cmd)
		}
		//fmt.Println("pos,face", currPos, currFace)

		if distance := currPos[0]*currPos[0] + currPos[1]*currPos[1]; distance > maxDistance {
			maxDistance = distance
		}

	}
	return maxDistance

}
func robotSim(commands []int, obstacles [][]int) int {
	//X:Y1,Y2...
	obMap := make(map[int]map[int]struct{})

	for _, pos := range obstacles {
		_, ok := obMap[pos[0]]
		if !ok {
			obMap[pos[0]] = make(map[int]struct{})
		}
		obMap[pos[0]][pos[1]] = struct{}{}
	}
	direction := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	currFace := 0
	currPos := []int{0, 0}
	tempPos := []int{0, 0}

	move := func(steps int) {
		for i := 0; i < steps; i++ {
			copy(tempPos, currPos)
			tempPos[0] += direction[currFace][0]
			tempPos[1] += direction[currFace][1]

			obSet := obMap[tempPos[0]]
			if obSet == nil {
				goto SetPos
			}
			if _, ok := obSet[tempPos[1]]; ok {
				break
			}
		SetPos:
			copy(currPos, tempPos)
		}
	}
	turn := func(cmd int) {
		if cmd == -1 {
			currFace = (currFace + 1) % 4
		} else if cmd == -2 {
			currFace = (currFace + 3) % 4
		}
	}
	maxDistance := math.MinInt64
	for _, cmd := range commands {
		if cmd < 0 {
			turn(cmd)
		} else {
			move(cmd)
		}
		//fmt.Println("pos,face", currPos, currFace)

		if distance := currPos[0]*currPos[0] + currPos[1]*currPos[1]; distance > maxDistance {
			maxDistance = distance
		}

	}
	return maxDistance
}
