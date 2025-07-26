package rwmutex

import (
	"fmt"
	"sync"
	"time"
)

// BankAccount 表示一个银行账户，使用RWMutex保护
type BankAccount struct {
	balance int
	rwMutex sync.RWMutex
}

// Deposit 存款操作（写操作）
func (ba *BankAccount) Deposit(amount int) {
	ba.rwMutex.Lock()         // 获取写锁
	defer ba.rwMutex.Unlock() // 确保释放锁

	fmt.Printf("准备存款: %d\n", amount)
	ba.balance += amount
	time.Sleep(100 * time.Millisecond) // 模拟耗时操作
	fmt.Printf("完成存款: %d, 新余额: %d\n", amount, ba.balance)
}

// Balance 查询余额（读操作）
func (ba *BankAccount) Balance() int {
	ba.rwMutex.RLock()         // 获取读锁
	defer ba.rwMutex.RUnlock() // 确保释放锁

	fmt.Println("正在查询余额...")
	time.Sleep(50 * time.Millisecond) // 模拟耗时操作
	return ba.balance
}

// TransactionLog 模拟交易日志查询（只读操作）
func (ba *BankAccount) TransactionLog() {
	ba.rwMutex.RLock()         // 获取读锁
	defer ba.rwMutex.RUnlock() // 确保释放锁

	fmt.Println("正在生成交易日志...")
	time.Sleep(80 * time.Millisecond) // 模拟耗时操作
	fmt.Printf("当前余额: %d\n", ba.balance)
}

func BankTest() {
	account := BankAccount{balance: 100}
	var wg sync.WaitGroup

	// 启动多个并发操作
	wg.Add(10)
	
	// 并发读取操作（可以同时进行）
	go func() {
		defer wg.Done()
		fmt.Println("余额查询结果:", account.Balance())
	}()
	
	go func() {
		defer wg.Done()
		account.TransactionLog()
	}()
	
	go func() {
		defer wg.Done()
		fmt.Println("余额查询结果:", account.Balance())
	}()

	// 并发写入操作（互斥进行）
	go func() {
		defer wg.Done()
		account.Deposit(50)
	}()
	
	go func() {
		defer wg.Done()
		account.Deposit(-20)
	}()
	

	go func() {
		defer wg.Done()
		fmt.Println("余额查询结果:", account.Balance())
	}()

	go func() {
		defer wg.Done()
		fmt.Println("余额查询结果:", account.Balance())
	}()

	go func() {
		defer wg.Done()
		fmt.Println("余额查询结果:", account.Balance())
	}()

	go func() {
		defer wg.Done()
		fmt.Println("余额查询结果:", account.Balance())
	}()

	go func() {
		defer wg.Done()
		fmt.Println("余额查询结果:", account.Balance())
	}()

	wg.Wait()
	fmt.Println("所有操作完成，最终余额:", account.Balance())
}