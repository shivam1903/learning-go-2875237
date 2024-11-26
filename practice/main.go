package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10, "Haww"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}
