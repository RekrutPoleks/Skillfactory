package main

import "fmt"

func main(){
	ch1 := make(chan int)
	go func(){
		for i:=0; i <= 100; i++{
			ch1 <- i
		}
		close(ch1)
	}()

	for val, ok := <- ch1; ok; val, ok = <- ch1{
		fmt.Print(val, " ")
	}
}
