package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("The detailed struct is %+v\n", poodle)
	fmt.Printf("Breed %v, Wieght %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 12
	fmt.Printf("Breed %v, Wieght %v\n", poodle.Breed, poodle.Weight)
	

}

// Dog is a struct
type Dog struct {
	Breed string
	Weight int

}
