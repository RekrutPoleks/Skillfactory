package main



import (
		"fmt"
		"sync"
		"time"
		)

var mu = sync.Mutex{}

func main() {
	var result uint64
	go intervalSum(&result, 0, 50000000)
	go intervalSum(&result, 50000000, 100000000)
	time.Sleep(time.Second) // give goroutines time to finish

	fmt.Println(result)

	otherResult := 0

	for i := 0; i < 100000000; i++ {
		otherResult += i
	}
	fmt.Println(otherResult)
}

func intervalSum(destination *uint64, start, end uint64) {
	for i := start; i < end; i++ {
		mu.Lock()
		result := *destination
		result += i
		*destination = result
		mu.Unlock()
	}
}