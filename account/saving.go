package account

import "github.com/wandersonsantos01/golang/client"

type Saving struct {
	Holder                     client.Holder
	Agency, Account, Operation int
	balance                    float64
}

func (acc *Saving) Withdraw(withdrawValue float64) bool {
	canWithdraw := withdrawValue > 0 && withdrawValue <= acc.balance
	if canWithdraw {
		acc.balance -= withdrawValue
		return true
	}
	return false
}

func (acc *Saving) Deposit(depositValue float64) (bool, float64) {
	if depositValue < 0 {
		return false, 0
	}
	acc.balance += depositValue
	return true, acc.balance
}

func (acc *Saving) Transfer(transferValue float64, toChecking *Checking) (bool, float64) {
	withdraw := acc.Withdraw(transferValue)
	if withdraw {
		deposit, _ := toChecking.Deposit(transferValue)
		if deposit {
			return true, acc.balance
		}
	}
	return false, 0
}

func (acc *Saving) GetBalance() float64 {
	return acc.balance
}
