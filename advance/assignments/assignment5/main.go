package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.balance += amount
}

func main() {
	account := &BankAccount{
		balance: 1000,
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			account.Deposit(10)
		}()
	}

	wg.Wait()

	fmt.Println("Final balance:", account.balance)
}
