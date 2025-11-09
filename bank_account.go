package main

import (
	"fmt"
	"sync"
)

type Account struct {
	Balance float32
	mu      sync.Mutex
}

func newAccount() *Account {
	return &Account{Balance: 0}
}

func (ac *Account) Deposit(amount float32) {
	ac.mu.Lock()
	ac.Balance += amount
	ac.mu.Unlock()
}

func (ac *Account) WithDraw(amount float32) {
	ac.mu.Lock()
	ac.Balance -= amount
	ac.mu.Unlock()
}

func main() {
	numTransactions := 10
	account := newAccount()

	var wg sync.WaitGroup
	for i := 0; i < numTransactions; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			account.Deposit(10)
		}()

		go func() {
			defer wg.Done()
			account.WithDraw(5)
		}()
	}

	wg.Wait()
	fmt.Println("Balance should be 5 * 10 = 50")
	fmt.Printf("Final balance %v", account.Balance)
}
