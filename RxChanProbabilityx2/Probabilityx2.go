package main

import (
	"fmt"
	"time"
)

func main() {
	tiker1 := time.NewTicker(time.Second)
	tiker2 := time.NewTicker(2 * time.Second)

	for {

		select {
		case <-tiker1.C:
			fmt.Println("Принято сообщение из первого канала")
		case <-tiker2.C:
			fmt.Println("Принято сообщение из второго канала")

		}
	}
}
