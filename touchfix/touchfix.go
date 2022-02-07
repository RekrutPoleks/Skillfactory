package main

import (
	"fmt"
)

func main() {
	// y = ax + b
	a1, b1, a2, b2 := inputK()
	if a1 == a2 {
		fmt.Println("Линии не имеют точки пересечения")
		return
	}
	x := (b2 - b1) / (a1 - a2)
	y := a1*x + b1
	fmt.Printf("X: %.2f, Y: %.2f", x, y)

}

func inputK() (float32, float32, float32, float32) {
	var a1, b1, a2, b2 float32
	fmt.Println("Введите коофициент a1")
	fmt.Scanln(&a1)
	fmt.Println("Введите коофициент b1")
	fmt.Scanln(&b1)
	fmt.Println("Введите коофициент a2")
	fmt.Scanln(&a2)
	fmt.Println("Введите коофициент b2")
	fmt.Scanln(&b2)
	return a1, b1, a2, b2
}
