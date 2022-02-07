package ReadConsole

import (
	"bufio"
	"fmt"
	"os"
)

func ReadCon() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("output: ", text)
}
