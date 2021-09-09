package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	owner   string
	balance int
	mutex   sync.Mutex
}

func (acc *BankAccount) deposit(value int) {
	acc.mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, acc.balance)
	acc.balance += value
	acc.mutex.Unlock()
}

func (acc *BankAccount) withdraw(value int) {
	acc.mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, acc.balance)
	acc.balance -= value
	acc.mutex.Unlock()
}

func demoAccount() {
	acc := BankAccount{
		owner:   "Cường",
		balance: 600,
	}
	go acc.withdraw(700)
	go acc.deposit(500)

	fmt.Printf("New Balance %d\n", acc.balance)
}
