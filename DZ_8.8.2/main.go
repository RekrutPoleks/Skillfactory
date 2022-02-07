package main

import (
	"fmt"

	"Skillfactory/DZ_8.8.2/autom"
)

func main() {

	// BMW, Mercedes и Dodge
	BMW := autom.NewAutoEuro("BMW", "I8", 260, 420, 180.0, 220.0, 110.0)
	Mercedes := autom.NewAutoEuro("Mercedes", "CLS", 220, 450, 190.0, 500.0, 120.0)
	Dodge := autom.NewAutoAmerica("Dodge", "RAM", 200, 600, 200.0, 600.0, 160.0)
	PrintCAR(&BMW)
	PrintCAR(&Mercedes)
	PrintCAR(&Dodge)
}

func PrintCAR(intauto autom.Auto) {
	auto := intauto.(*autom.Automobil)
	template := "Brand: %s\n\tModel: %s\n\tMaxSpeed: %d\n\tPower: %d\n\tSize:\n\t\tlength: %.1f\n\t\tWidth: %.1f\n\t\tHeight: %.1f\n\n"
	a, b, c := auto.Getsize()
	fmt.Printf(template, auto.Brand(), auto.Model(), auto.MaxSpeed(), auto.EnginePower(), a, b, c)

}
