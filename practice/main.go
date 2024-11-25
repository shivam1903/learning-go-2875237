package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 5,6,7
	intSum := i1 + i2+ i3
	fmt.Println("The Integer sum is  = ", intSum)

	f1, f2, f3 := 12.5, 18.7, 123.567
	var floatSum = f1 + f2 + f3
	fmt.Println("This is the Float Sum ", floatSum)
	fmt.Println("Math")

	floatSum = math.Round(floatSum*100)/100
	fmt.Println("Updated Float Sum ", floatSum)

	circleRad := 15.5
	circumference := circleRad * 2 * math.Pi
	fmt.Printf("The circumference is %.1f\n", circumference)

}
