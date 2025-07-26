package main

import (
	"fmt"
	"sync"

	rwm "example.com/go-demo/ch9/rwmutex"
)

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
	fmt.Printf("deposit %d successfully\n", amount)
}

func Balance() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("now balance is ", balance)
}

func mutexTest() {
    Deposit(20)
    Deposit(10)
    Balance()
    Deposit(30)
    Balance()
}

func main() {
	mutexTest()
	fmt.Println("All operations completed. Final balance:", balance)
	rwm.BankTest()
}