package tests

import (
	"fmt"

	"github.com/renatospaka/design-pattern-go/behavioral/command"
)

func TestCommand() {
	fmt.Println("Command -> raw")

	bankAcc := command.BankAccount{}
	deposit1 := command.NewBankAccountCommand(&bankAcc, command.Deposit, 100)
	deposit1.Call()
	fmt.Println(bankAcc)
	deposit2 := command.NewBankAccountCommand(&bankAcc, command.Withdraw, 50)
	deposit2.Call()
	fmt.Println(bankAcc)
}
