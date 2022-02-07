package intRingV2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	const (
		sizeBuffer = 1000
		timeinterbefore = 5
		durationTest = 5
		countGoro = 10
	)
	exit, start  := 0, 0
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().Unix())
	ringBuf := NewRing(sizeBuffer)
	chanExit := make(chan int)
	goroWrite := func (){
		tiker := time.NewTicker(time.Duration(1+rand.Intn(timeinterbefore-1))*time.Second)
		start++
		println("start: ", start)
		defer wg.Done()
		for {
			select {
			case <-tiker.C:
				val := rand.Intn(1000)
				err := ringBuf.Write(val)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Write: ", val)
				}
			case <-chanExit:
				println("Exit", exit)
				return

			}


		}
	}
	goroRead := func(){
		tiker := time.NewTicker(time.Duration(1+rand.Intn(timeinterbefore-1))*time.Second)
		start++
		println("start: ", start)
		defer wg.Done()
		for {
			select {
			case <- tiker.C:
				val, err := ringBuf.Read()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Read:", val)
				}
			case <-chanExit:
				println("Exit", exit)
				return
			}


		}
	}


	for i:=0; i < countGoro; i++{
		wg.Add(2)
		go goroRead()
		go goroWrite()
	}
	go func () {
		<- time.After(durationTest * time.Second)
		close(chanExit)
	}()
	wg.Wait()
	fmt.Println("TEST OK, Good bye!!!!!!!")
}