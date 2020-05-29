package main

import (
	"fmt"
	"math/rand"
)

func main() {
	command1 := "walk outside"
	if command1 == "walk outside" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	command2 := "walk inside"
	switch command2 {
	case "walk outside":
		fmt.Println("walk outside")
		fmt.Println("walk outside")
	case "walk inside":
		fmt.Println("walk inside")
		fmt.Println("walk inside")
	default:
		fmt.Println("walk default")
	}

	count := 10
	for count > 0 {
		fmt.Println(count)
		count--
	}
	fmt.Println("Liftoff!")

	degrees := 0
	for {
		fmt.Println(degrees)
		degrees++
		if degrees >= 20 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}
