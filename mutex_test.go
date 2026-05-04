package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Mutex untuk menghadle race condition
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}

type BankAccount struct {
	RWMutex sync.Mutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	defer account.RWMutex.Unlock()
	account.Balance += amount
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.Lock()
	defer account.RWMutex.Unlock()
	return account.Balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			account.AddBalance(1)
			fmt.Println(account.GetBalance())
		}
	}
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	defer user1.Unlock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	defer user2.Unlock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "pai",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "aminu",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)

	fmt.Println("User", user1.Name, "Balance", user1.Balance)
	fmt.Println("User", user2.Name, "Balance", user2.Balance)
}
