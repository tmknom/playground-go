package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	year := 2020
	fmt.Printf("%T 型：%v\n", year, year)
	fmt.Printf("%T 型：%[1]v\n", year)

	var red uint8 = 255
	red++
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number)

	var green uint8 = 3
	fmt.Printf("%08b\n", green)

	green++
	fmt.Printf("%08b\n", green)

	fmt.Println(math.MaxInt8)

	future := time.Unix(12622780800, 0)
	fmt.Println(future)
}
