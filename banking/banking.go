package banking

import (
	"errors"
	"fmt"
)

// Account struct
type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("not enough money")

// NewAccount account constructor
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount from account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a account) String() string {
	return fmt.Sprint(a.owner, " has ", a.balance)
}
