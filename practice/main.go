package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter number: ")
	inputInt, _ := reader.ReadString('\n')
	numFloat, err := strconv.ParseFloat(strings.TrimSpace(inputInt), 64)
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("Given number = ", numFloat)
	}


}
