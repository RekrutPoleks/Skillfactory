package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main(){
	 var count int32 = 0
	 var wg sync.WaitGroup
	 adder := func(){
		 defer wg.Done()
		 atomic.AddInt32(&count, 1)
	 }
	 for i:= 10; i > 0; i--{
		 wg.Add(1)
		 go adder()
	 }
	 wg.Wait()
	 fmt.Println("count = ", count)
}
