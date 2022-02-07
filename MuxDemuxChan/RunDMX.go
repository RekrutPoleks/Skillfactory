package main

import (
	"Skillfactory/MuxDemuxChan/structurDMX"
	"fmt"
	"sync"
)

func main(){
	const countGorutine = 40
	dmx := structurDMX.NewDMXUnlimit()
	dmxClose := structurDMX.NewDMXUnlimit()
	wg := sync.WaitGroup{}
	inDmxChan := dmx.GetInChanDMX()
	exitChan := dmxClose.GetInChanDMX()

	printGorutineChanExit := func(number int, ch, chExit chan structurDMX.Message){
		defer wg.Done()
		for {
			select {
			case msg := <-ch:
				fmt.Printf("Gorutine №%d, принято сообщение %s\n", number, msg)
			case <-chExit:
				fmt.Printf("Gorutine №%d, остановлена\n", number)
				return
			}
		}
	}
	wg.Add(countGorutine)
	for i:=0; i< countGorutine; i++{
		ch := make(chan structurDMX.Message)
		chEx := make(chan structurDMX.Message)
		_ = dmx.AddChan(ch)
		_ = dmxClose.AddChan(chEx)
		go printGorutineChanExit(i, ch, chEx)
	}
	go dmx.RunDMXOnce()
	go dmxClose.RunDMXOnce()
	inDmxChan <- "Happy New Year"
	exitChan <- ""
	wg.Wait()
	fmt.Println("Good Bye")
}

