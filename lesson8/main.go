package main

import (
	"fmt"
	"math/big"
)

func main() {
	var distance int64 = 41.3e12
	fmt.Println(distance)

	days0 := distance / 299792 / 86400
	fmt.Println(days0)

	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance2 := new(big.Int)
	distance2.SetString("2400000000000000000000", 10)
	fmt.Println(distance2)

	seconds := new(big.Int)
	seconds.Div(distance2, lightSpeed)
	fmt.Println(seconds)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println(days)
}
