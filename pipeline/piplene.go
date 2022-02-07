package main

import (
	"Skillfactory/intRingV2.0"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main(){
	const (
		lifeTime = 20 // секунд
		rangeReadTime = 10 // секунд
		countBuffer = 1000
	)

	buf := intRingV2.NewRing(countBuffer)
	chClose := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	controlTimeLife := func(){
		defer wg.Done()
		<- time.After(lifeTime*time.Second)
		close(chClose)
	}

	chInput := inputData(chClose)
	go controlTimeLife()
	go writeBuffer(buf, filterNotМultipleThree(filterNotNegativeAndZero(chInput, chClose), chClose), chClose)
	go readBuffer(buf, rangeReadTime, chClose)
	wg.Wait()

}

func inputData(chClose chan bool) <-chan *int{
	chInPip:= make(chan *int)
	go func() {
		for {
			select {
				case <-chClose:
					return
				default:
					var inputVal string
					_, err := fmt.Scanln(&inputVal)
					if err != nil {
						fmt.Println(err)
						continue
					}
					val, err := strconv.Atoi(inputVal)
					if err != nil {
						fmt.Println("Введены не верные данные:", inputVal)
					} else {
						chInPip <- &val
					}
			}
		}
	}()
	return chInPip
}

func filterNotМultipleThree (chIn <-chan *int, chClose chan bool) <-chan *int{
	chfilter := make(chan *int)
	go func(){
		for {
			select {
			case val := <-chIn:
				if *val%3 != 0 {
					chfilter <- val
				}
			case <-chClose:
				return
			}
		}
	}()
	return chfilter
}

func filterNotNegativeAndZero(chIn <-chan *int, chClose chan bool) <-chan *int{
	chfilter := make(chan *int)
	go func(){
		for {
			select {
			case val := <-chIn:
				if *val > 0 {
					chfilter <- val
				}
			case <-chClose:
				return
			}
		}
	}()
	return chfilter
}

func writeBuffer(buf *intRingV2.Ring, chIn <-chan *int, chClose chan bool){
	for {
		select {
		case val := <-chIn:
			_ = buf.Write(*val)
		case <- chClose:
			return
		}
	}
}

func readBuffer (buf *intRingV2.Ring, rangeTime int, chClose chan bool){
	tiker := time.NewTicker(time.Duration(rangeTime)*time.Second)
	for {
		select {
		case <-tiker.C:
			fmt.Println("Data from the buffer: ")
			for {
				val, err := buf.Read()
				if err != nil{
					fmt.Println()
					fmt.Println("Data read, buffer empty")
					break
				} else {
					fmt.Printf("%d ", val)
				}
			}
		case <-chClose:
			return
		}
	}

}