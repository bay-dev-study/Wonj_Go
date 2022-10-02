package main

import (
	"fmt"
	"std/github.com/wonj1012/learngo/banking"
)

func main() {
	account := banking.NewAccount("wonj")
	account.Deposit(100)
	fmt.Println(account)
	account.Withdraw(50)
	fmt.Println(account)
}
