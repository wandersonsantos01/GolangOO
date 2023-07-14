package account

import "github.com/wandersonsantos01/golang/client"

type Checking struct {
	Holder  client.Holder
	Agency  int
	Checking int
	balance float64
}

func (acc *Checking) Withdraw(withdrawValue float64) bool {
	canWithdraw := withdrawValue > 0 && withdrawValue <= acc.balance
	if canWithdraw {
		acc.balance -= withdrawValue
		return true
	}
	return false
}

func (acc *Checking) Deposit(depositValue float64) (bool, float64) {
	if depositValue < 0 {
		return false, 0
	}
	acc.balance += depositValue
	return true, acc.balance
}

func (acc *Checking) Transfer(transferValue float64, toChecking *Checking) (bool, float64) {
	withdraw := acc.Withdraw(transferValue)
	if withdraw {
		deposit, _ := toChecking.Deposit(transferValue)
		if deposit {
			return true, acc.balance
		}
	}
	return false, 0
}

func (acc *Checking) GetBalance() float64 {
	return acc.balance
}
