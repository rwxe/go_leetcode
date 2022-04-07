package leetcode

func winnerOfGame(colors string) bool {
	counter := 1
	ACounter, BCounter := 0, 0
	prevColor := colors[0]
	for i := 1; i < len(colors); i++ {
		if prevColor == colors[i] {
			counter++
		} else {
			if v := counter - 2; v > 0 {
				if prevColor == 'A' {
					ACounter += v
				} else {
					BCounter += v
				}
			}
			counter = 1
		}
		prevColor = colors[i]
	}
	if counter != 1 {
		if v := counter - 2; v > 0 {
			if prevColor == 'A' {
				ACounter += v
			} else {
				BCounter += v
			}
		}
	}
	if ACounter > BCounter {
		return true
	} else {
		return true

	}
}
