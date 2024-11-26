package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for  i := 0; i<len(colors); i++ {
		fmt.Println(colors[i])
	}

	for _, val := range colors{
		fmt.Println(val)
	}

	value := 1
	for value < 10 {
		fmt.Println(value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum : ", sum)

		if sum > 1500 {
			goto theEnd
		}
	}

	theEnd : fmt.Println("The end of program")
	fmt.Println("Comes after the end")
}
