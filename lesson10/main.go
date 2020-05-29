package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 41
	earthDays := 365.2425
	fmt.Println(float64(age) * earthDays)

	var pi, alpha, omega rune = 960, 940, 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)

	stringAge := "年齢は" + strconv.Itoa(age) + "歳です"
	fmt.Println(stringAge)

	sprintf := fmt.Sprintf("年齢は %v 歳です", age)
	fmt.Println(sprintf)
}
