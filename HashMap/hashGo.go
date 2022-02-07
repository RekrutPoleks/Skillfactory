package main

import (
	"Skillfactory/HashMap/HashMapa"
	"Skillfactory/HashMap/MyHashFunc"
	"fmt"
)

func main(){
	fmt.Println(MyHashFunc.HashStr("dskfks"))
	fmt.Println(MyHashFunc.HashStr("dskfks 23 1 34fff"))
	fmt.Println(MyHashFunc.HashStr("lprijr"))
	fmt.Println(MyHashFunc.HashStr("abc"))
	fmt.Println(MyHashFunc.HashStr("cba"))
	mapa := HashMapa.NewHashMap()
	mapa.Set("abc", "ertyuuiio")
	fmt.Println(mapa.Get("abc"))
	mapa.Set("qwerty", "kijhygtrf")
	fmt.Println(mapa)
	mapa.Delete("abc")
	fmt.Println(mapa)
	fmt.Println(mapa.Get("abc"))

}

