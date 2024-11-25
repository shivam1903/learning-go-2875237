package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)

	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", input)

}
