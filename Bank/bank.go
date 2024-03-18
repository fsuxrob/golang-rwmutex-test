package Bank

import (
	"fmt"
	"sync"
)

type BankClient interface {
	Deposit(amount int)
	Withdraw(amount int) error
	Balance() int
}

type BankAccount struct {
	bMutex  sync.RWMutex
	balance int
}

func NewBankAccount(balance int) *BankAccount {
	return &BankAccount{balance: balance}
}

func (b *BankAccount) Deposit(amount int) {
	b.bMutex.Lock()
	b.balance += amount
	b.bMutex.Unlock()
}

func (b *BankAccount) Withdraw(amount int) error {
	b.bMutex.Lock()
	defer b.bMutex.Unlock()
	if b.balance < amount {
		return fmt.Errorf("balance %d less than amount %d", b.balance, amount)
	}
	b.balance -= amount
	return nil
}

func (b *BankAccount) Balance() int {
	b.bMutex.RLock()
	defer b.bMutex.RUnlock()
	return b.balance
}
