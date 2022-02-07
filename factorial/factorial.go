package main

import (
	"fmt"
)

func main() {
	// var fact uint64
	var input uint64
	fmt.Scan(&input)
	fmt.Println(rekursia(input, 1))
}

func rekursia(input uint64, step uint64) uint64 {

	if step == input {
		return step
	} else {
		fact := step * rekursia(input, step+1)
		fmt.Println(fact)
		return fact
	}
}

func obratRekursia(input uint64) uint64 {
	if input < 2 {
		return 1
	}
	return input * obratRekursia(input-1)
}
