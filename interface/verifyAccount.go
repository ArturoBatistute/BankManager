package verifyAccountInterface

type VerifyAccount interface {
	Withdraw(value float64)
}

func PayBill(account VerifyAccount, billValue float64) {
	account.Withdraw(billValue)
}
