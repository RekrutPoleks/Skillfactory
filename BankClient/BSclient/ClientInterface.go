package BSclient

import (
	"fmt"
	"sync"
)

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}


type BClient struct {
	depositRWMutex sync.RWMutex
	deposit int
}

func NewBClient() *BClient {
	return &BClient{sync.RWMutex{}, 0}
}

func (bc *BClient) Deposit(amount int){
	bc.depositRWMutex.Lock()
	defer bc.depositRWMutex.Unlock()
	bc.deposit+=amount
}

func (bc *BClient) Balance() int {
	bc.depositRWMutex.RLock()
	defer bc.depositRWMutex.RUnlock()
	return bc.deposit
}

func (bc *BClient) Withdrawal(amount int) error{
	bc.depositRWMutex.Lock()
	defer bc.depositRWMutex.Unlock()
	if amount > bc.deposit{
		return fmt.Errorf("Your balance %d is less than the specified amount", bc.deposit)
	}
	bc.deposit -= amount
	//fmt.Println("Withdrawn from deposit ", amount)
	//fmt.Println("Your balance ", bc.deposit)
	return nil
}

