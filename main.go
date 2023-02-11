package main

import (
	"BankManager/account"
	i "BankManager/interface"
	"BankManager/user"
	"fmt"
)

func main() {

	johnClient := user.Client{Name: "John", ID: 00, Occupation: "Engineer"}
	stephanClient := user.Client{Name: "Stephan", ID: 01, Occupation: "Doctor"}

	johnAccount := account.CheckingAccount{
		Client: johnClient,
		Agency: 001,
		Number: 2931}

	stephanAccount := account.CheckingAccount{
		Client: stephanClient,
		Agency: 001,
		Number: 2591}

	johnAccount.Deposit(100)
	stephanAccount.Deposit(50)

	err := johnAccount.TransferBetweenAccounts(&stephanAccount, 20)
	if err != nil {
		fmt.Println(err)
	}

	i.PayBill(&johnAccount, 10)

	fmt.Println(johnAccount)
	fmt.Println(stephanAccount)
}
