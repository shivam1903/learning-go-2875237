package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var res string

	switch dow{
	case 1:
		res = "It is Sunday"
		// fallthrough
		// fallthrough is used when we want to move to the next case, just like skip
	
	case 2:
		res = "It's monday"
	
	default:
		res = "It is a weekday"

	}

	fmt.Println(res)
}
