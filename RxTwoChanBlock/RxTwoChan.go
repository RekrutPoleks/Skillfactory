package main

import (
	"fmt"
	"math/rand"
	"time"
)

func  main(){
	ch1 := retRandInt()
	ch2 := retRandInt()
	for ;(ch1 != nil)&&(ch2 != nil) ;{
		select {
		case n, ok :=<-ch1:
			if !ok {
				ch1 = nil
			}
			fmt.Println("Return ch1 value int= ", n)
		case n, ok :=<- ch2:
			if !ok {
				ch2 = nil
			}
			fmt.Println("Return ch2 value int=", n)
		default:
			fmt.Println(time.Now().Format(time.RFC822))
			time.Sleep(time.Second)
		}

	}
	fmt.Println("Good Bye!!!!!!")
}

func retRandInt()  <- chan int{
	ch := make(chan int)
	const countIter int = 10
	delay := 1 + rand.Intn(4)
	go func(){
		for i:= countIter; i >0; i--{
			time.Sleep(time.Duration(delay)*time.Second)
			ch <- rand.Intn(999)
		}
		close(ch)
	}()
	return ch
}