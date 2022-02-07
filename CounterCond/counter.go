package main

import (
	"fmt"
	"sync"
)

const step int = 1
const interationAmount int = 1000

func main() {
	var counter int = 0
	var c = sync.NewCond(&sync.Mutex{})
	increment := func() {
		c.L.Lock()
		counter += step
		if counter == interationAmount {
			c.Signal()
		}
		c.L.Unlock()
	}
	c.L.Lock()
	for i := interationAmount; i > 0; i-- {
		go increment()
	}
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
