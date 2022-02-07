package main

import "fmt"

func main(){
	inputValue := make(map[int]interface{})
	var countIN1, countIN2, input int
	resultArray := make([]int,0)
	cashe := make(map[int]interface{})
	fmt.Print("Enter first array size: ")
	fmt.Scanln(&countIN1)
	fmt.Print("Enter second array size: ")
	fmt.Scanln(&countIN2)
	fmt.Println("Enter first array: ")
	for i :=0 ; i < countIN1; i++{
		fmt.Scanln(&input)
		inputValue[input] = nil

	}
	fmt.Println("Enter second array: ")
	for i :=0 ; i < countIN2; i++{
		fmt.Scanln(&input)
		_, ok := inputValue[input];
		_, okCashe := cashe[input]
		if ok && !okCashe {
			resultArray = append(resultArray, input)
			cashe[input] = nil
		}

	}
	fmt.Println(resultArray)
}
