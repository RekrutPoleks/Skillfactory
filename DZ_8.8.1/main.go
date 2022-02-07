package main

import (
	"Skillfactory/DZ_8.8.1/electronic"
	"fmt"
)

func main() {
	huawey := electronic.NewAndroidPhone("Huawey", "Nova3")
	iphone7 := electronic.NewApplePhone("IPhone 7")
	stationT := electronic.NewRadioPhone("Sputnic", "Klever", 12)
	printCharacteristics(huawey)
	printCharacteristics(iphone7)
	printCharacteristics(stationT)

}

func printCharacteristics(Phone interface{}) {
	switch phone := Phone.(type) {
	case electronic.AndroidPhone:
		fmt.Printf("%s, %s, %s, %s\n", phone.Brand(), phone.Model(), phone.OS(), phone.Type())
	case electronic.ApplePhone:
		fmt.Printf("%s, %s, %s, %s\n", phone.Brand(), phone.Model(), phone.OS(), phone.Type())
	case electronic.RadioPhone:
		fmt.Printf("%s, %s, %d, %s\n", phone.Brand(), phone.Model(), phone.ButtonsCount(), phone.Type())
	default:
		fmt.Printf("%T\n", Phone)

	}
}
