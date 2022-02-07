package main

import (
	"Skillfactory/IntRing/StructurRing"
	"fmt"
	"math/rand"
	"sync"
	"time"

)

func main(){
	const (
		interval  = 1000
		countGorutine  = 60
		sizeBuffer          = 1000
		intervalStartBefore = 5
		timeLife = 1
		countWriter = 5
	)
	rand.Seed(time.Now().Unix())
	bufferRing := StructurRing.NewIntRing(sizeBuffer, countWriter)
	chanClose := make(chan bool)
	cond := sync.NewCond(&sync.Mutex{})

	go func(){
		<-time.After(timeLife * time.Second)
		close(chanClose)
		cond.Signal()
	}()
	creator := func(period int) {
		tiker := time.NewTicker(time.Duration(period) *time.Millisecond)
		for {
			select {
			case <-chanClose:
				return
			case <- tiker.C:
				el := rand.Intn(interval)
				err := bufferRing.Write(el)
				if err != nil{
					fmt.Println(err)
				}else{fmt.Println("write el: ", el)}

			}

		}
	}

	reader := func(period int){
		tiker := time.NewTicker(time.Duration(period)*time.Millisecond)
		for {
			select {
				case <- chanClose:
					return
				case <- tiker.C:
					el:= bufferRing.Read()
					fmt.Println("read el:", el)
			}
		}
	}
	cond.L.Lock()
	for i:= 0 ; i < countGorutine; i ++{
		go creator(1+rand.Intn(intervalStartBefore-1))
		go reader(1+rand.Intn(intervalStartBefore-1))
	}
	cond.Wait()
	cond.L.Unlock()

}
