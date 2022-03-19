package leetcode

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	//ac1->ac2
	account1, account2 = account1-1, account2-1
	if 0 <= account1 && 0 <= account2 && account1 < len(this.balance) && account2 < len(this.balance) {
		if money <= this.balance[account1] {
			this.balance[account1] -= money
			this.balance[account2] += money
			return true
		}
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	account = account - 1
	if 0 <= account && account < len(this.balance) {
		this.balance[account] += money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	account = account - 1
	if 0 <= account && account < len(this.balance) {
		if this.balance[account] >= money {
			this.balance[account] -= money
			return true
		}
	}
	return false

}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
