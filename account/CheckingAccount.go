package account

import (
	"BankManager/user"
	"errors"
	"fmt"
)

type CheckingAccount struct {
	Client  user.Client
	Agency  int
	Number  int
	balance float64
}

func (account *CheckingAccount) Withdraw(valueToWithdraw float64) {

	if valueToWithdraw > account.balance || valueToWithdraw < 0 {
		fmt.Println("Error, operation not authorized.")
	} else {
		account.balance -= valueToWithdraw
		fmt.Println("Succesfully withdraw.")
	}
}

func (account *CheckingAccount) Deposit(valueToDeposit float64) {

	if valueToDeposit < 0 {
		fmt.Println("Error, operation not authorized.")
	} else {
		account.balance += valueToDeposit
		fmt.Println("Succesfully deposited.")
	}
}

func (sourceAccount *CheckingAccount) TransferBetweenAccounts(receiverAccount *CheckingAccount, valueToTransfer float64) error {

	if valueToTransfer > sourceAccount.balance {
		return errors.New("error, value to transfer must be lower than source account balance")
	}

	sourceAccount.Withdraw(valueToTransfer)
	receiverAccount.Deposit(valueToTransfer)

	fmt.Println("Transfer between account successfully.")
	return nil
}

func (account *CheckingAccount) ShowBalance() float64 {

	return account.balance
}
