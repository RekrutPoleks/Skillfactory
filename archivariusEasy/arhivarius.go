package main

import "fmt"

func main() {
	var input string = "aaassssdffffg"
	var output string
	var count int
	var cashe rune = []rune(input)[0]
	for _, char := range input {
		if char == cashe {
			count++
			continue
		} else {
			output = output + fmt.Sprintf("%c%d", cashe, count)
			cashe = char
			count = 1
		}
	}
	output += fmt.Sprintf("%c%d", cashe, count)
	fmt.Println(output)
}
