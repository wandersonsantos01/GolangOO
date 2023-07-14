package main

import (
	"fmt"

	"github.com/wandersonsantos01/golang/account"
)

func main() {
	larsAccount := account.Saving{}
	larsAccount.Deposit(200)
	account.PayBillet(&larsAccount, 75)

	fmt.Println(larsAccount)

	trujiloAccount := account.Checking{}
	trujiloAccount.Deposit(100)
	account.PayBillet(&trujiloAccount, 99)

	fmt.Println(trujiloAccount)

	// kirkAccount := account.Checking{}
	// kirkAccount.Deposit(20)

	// fmt.Println(kirkAccount.GetBalance())

	// jamesClient := client.Holder{Name: "James", Document: "12345678900", Occupation: "Musician"}
	// jamesAccount := account.Cheking{Holder: jamesClient, Agency: 123456, Balance: 900}

	// fmt.Println(jamesAccount)

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

	// // Withdraw
	// var georgeAccount *Account
	// georgeAccount = new(Account)
	// georgeAccount.holder = "George"
	// georgeAccount.balance = 400

	// fmt.Println(georgeAccount.Withdraw(200.))
	// fmt.Println(georgeAccount)

	// Deposit
	// var angusAccount *Account
	// angusAccount = new(Account)
	// angusAccount.holder = "Angus"
	// angusAccount.balance = 100

	// status, balance := angusAccount.Deposit(150.99)
	// fmt.Println("status=", status)
	// fmt.Println("balance=", balance)

	// axelAccount := account.Account{Holder: "Axel", Balance: 100}
	// slashAccount := account.Account{Holder: "Slash", Balance: 50}

	// fmt.Println(axelAccount)
	// fmt.Println(slashAccount)

	// axelAccount.Transfer(-20, &slashAccount)

	// fmt.Println(axelAccount)
	// fmt.Println(slashAccount)

}
