package main

import (
	"Skillfactory/GoroutineCounterSyncChan/Counter"
	"fmt"
)

func main(){
	counter := Counter.NewCounter()
	counter.Start()
	fmt.Println(counter.GetResult())
}
