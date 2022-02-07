package main

import (
	BankClient "Skillfactory/BankClient/BSclient"
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().Unix())
	BC := BankClient.NewBClient()
	addDeposit := func(){
		for {
			time.Sleep(time.Duration(rand.Intn(500)+500)*time.Millisecond)
			BC.Deposit(rand.Intn(10))
			//fmt.Println("add")
		}

	}
	delDeposit := func(){
		for {
			time.Sleep(time.Duration(rand.Intn(500)+500)*time.Millisecond)
			if err := BC.Withdrawal(rand.Intn(10)); err != nil {
				fmt.Println(err)
			}
			//fmt.Println("del")
		}
	}
	for i:=0; i <= 5; i++{
		go addDeposit()
		go addDeposit()
		go delDeposit()
	}
	var inp string
	for  {
		switch fmt.Scanln(&inp); inp{
		case "deposit":
			fmt.Print("enter the amount to deposit : ")
			var amount int
			fmt.Scanln(&amount)
			BC.Deposit(amount)
			fmt.Println("Your balance ", BC.Balance())
		case "balance":
			fmt.Println("Balance: ", BC.Balance())
		case "withdrawal":
			fmt.Print("enter the withdrawal amount: ")
			var amount int
			fmt.Scanln(&amount)
			if err := BC.Withdrawal(amount); err != nil{
				fmt.Println(err)
			}else{
				fmt.Println("Withdrawn from deposit ", amount)
				fmt.Println("Your balance ", BC.Balance())
			}
		case "exit":
			return
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
		}
	}
}
