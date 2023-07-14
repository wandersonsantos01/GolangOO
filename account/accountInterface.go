package account

func PayBillet(account checkAccount, billetValue float64) {
	account.Withdraw(billetValue)
}

type checkAccount interface {
	Withdraw(value float64) bool
}
