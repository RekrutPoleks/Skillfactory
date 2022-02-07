package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
workArray := [10]uint
for i:=0; len(workArray) < i; i++{
    fmt.Scan(&workArray[i])
        
}

var a,b int

for i:=1; i <= 3; i++{         
    fmt.Scan(&a, &b)
    workArray[a], workArray[b] = workArray[b], workArray[a]
}
for i:= 0; i < len(workArray); i++{
    fmt.Printf("%d ", workArray[i])
}
}

