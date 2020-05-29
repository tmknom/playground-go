package main

import (
	"fmt"
	"math/rand"
)

var global = "global"

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}

	if num := rand.Intn(10); num == 0 {
		fmt.Println("zero")
	} else if num == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("other")
	}

	fmt.Println(global)
}
