package main

import "fmt"

type Account struct {
	holder  string
	agency  int
	account int
	balance float64
}

func (c *Account) Withdraw(withdrawValue float64) bool {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance
	if canWithdraw {
		c.balance -= withdrawValue
		return true
	}
	return false
}

func main() {
	// paulAccount := Account{holder: "Paul", agency: 123, account: 456789, balance: 235.78}
	// fmt.Println(paulAccount)

	// paulAccountClone := Account{holder: "Paul", agency: 123, account: 456789, balance: 235.78}
	// fmt.Println(paulAccount == paulAccountClone)

	// ringoAccount := Account{"Ringo", 456, 789789, 364.7}
	// fmt.Println(ringoAccount)

	// ringoAccountClone := Account{"Ringo", 456, 789789, 364.7}
	// fmt.Println(ringoAccount == ringoAccountClone)

	// var johnAccount *Account
	// johnAccount = new(Account)
	// johnAccount.holder = "John"
	// johnAccount.balance = 100
	// fmt.Println(johnAccount)

	// var johnAccountClone *Account
	// johnAccountClone = new(Account)
	// johnAccountClone.holder = "John"
	// johnAccountClone.balance = 100
	// fmt.Println(*johnAccount == *johnAccountClone)

	// fmt.Println(&johnAccount, &johnAccountClone)

	var georgeAccount *Account
	georgeAccount = new(Account)
	georgeAccount.holder = "George"
	georgeAccount.balance = 400

	fmt.Println(georgeAccount.Withdraw(200.))
	fmt.Println(georgeAccount)

}
