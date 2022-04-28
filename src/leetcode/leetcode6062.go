package leetcode

//未完成
type ATM struct {
	//20,50,100,200,500
	cashs [][2]int
}

func Constructor6062() ATM {
	cashs := [][2]int{
		{20, 0},
		{50, 0},
		{100, 0},
		{200, 0},
		{500, 0},
	}
	return ATM{cashs: cashs}
}

func (atm *ATM) Deposit(banknotesCount []int) {
	for i, v := range banknotesCount {
		atm.cashs[i][1] += v
	}
}

func (atm *ATM) Withdraw(amount int) []int {
	bI := 0
	bN := 0
	for i := len(atm.cashs) - 1; i >= 0; i-- {
		if atm.cashs[i][1] > 0 {
			bI = i
			bN = atm.cashs[bI][0]
			break
		}
	}
	if amount < bN {
		return []int{-1}
	}
	return []int{-1}

}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
