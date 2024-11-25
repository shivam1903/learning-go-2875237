package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("The current time date is ", n)
	fmt.Println("Dates and times")

	t := time.Date(2009, time.November, 10, 10,10, 0, 0, time.UTC)
	fmt.Println("Random time created is ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 10:10:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)

}
