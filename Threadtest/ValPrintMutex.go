package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var countTime, value int
	fmt.Scan(&countTime)
	tk := tiktak{
		sync.Mutex{},
		value,
	}
	go tk.IncrementValue()
	go tk.PPrint()

	time.Sleep(time.Duration(countTime)*time.Second )

}

type tiktak struct {
	mutex sync.Mutex
	value int
}

func (tk *tiktak)PPrint(){
	for {
			time.Sleep(500*time.Millisecond)
			tk.mutex.Lock()
			fmt.Println(tk.value)
			tk.mutex.Unlock()

	}
}

func (tk *tiktak)IncrementValue(){
	for {
			time.Sleep(1000*time.Millisecond)
			tk.mutex.Lock()
			tk.value+=1
			tk.mutex.Unlock()
	}
}