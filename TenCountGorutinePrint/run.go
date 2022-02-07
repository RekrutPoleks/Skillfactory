package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	const (
		countPrint int = 10
		countGorutine int = 5
	)
	wg := sync.WaitGroup{}
	printGo := func(number int){
		defer wg.Done()
		for i := countPrint; i > 0; i--{
			fmt.Println("Gorutine name ", number)
			time.Sleep(5 *time.Millisecond)
		}
	}
	wg.Add(countGorutine)
	for i:= countGorutine; i > 0; i--{
		go printGo(i)
	}
	wg.Wait()
}

