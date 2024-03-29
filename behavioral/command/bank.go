package command

import "fmt"

type Banking interface {
	Deposit(amount int)
	Withdraw(amount int) bool
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

var overDraftLimit = -500
type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) *BankAccount {
	return &BankAccount{
		balance: balance,
	}
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance - amount >= overDraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)	
		return true	
	}
	return false
}
