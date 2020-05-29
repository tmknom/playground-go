package main

import (
	"fmt"
)

func main() {
	days := 365.2425
	fmt.Printf("%.2f\n", days)

	var weeklyDays float64 = 7
	fmt.Println(weeklyDays)

	var zero float64
	fmt.Println(zero)

	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)
}
