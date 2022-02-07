package main

import (
	"fmt"
	"strconv"
)

func main() {
	// readSringtoConvert()
	ReadInteger()
}

func readSringtoConvert() {
	var text string
	fmt.Print("input number: ")

	_, err := fmt.Scanln(&text)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось прочитать строку: %s", err))
	}

	number, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось сконвектировать строку %s", err))

	}
	fmt.Println(number)
}

func ReadInteger() {
	var itext int
	fmt.Println("Введите число: ")
	_, err := fmt.Scanln(&itext)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось прочитать строку %s", err))
	}

	fmt.Println(itext)
}
