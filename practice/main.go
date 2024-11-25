package main

import (
	"fmt"
)

func main() {

	var colors [3]string

	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"

	fmt.Println("This is the array ", colors)

	var number = [5]int64{1,2,3,4,5}
	fmt.Println("this is the number array ", number)

	fmt.Println("Number of numbers ", len(number))

}
