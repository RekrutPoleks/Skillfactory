package main

import (
	"Skillfactory/calcStruct/calc"
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string
	fmt.Println("Введите число")
	fmt.Scan(&num1)
	fmt.Println("Введите оператор")
	fmt.Scan(&operator)
	fmt.Println("Введите число")
	fmt.Scan(&num2)
	calculat := calc.NewCalculator()
	if answer, err := calculat.Calculate(num1, num2, operator); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}

}
