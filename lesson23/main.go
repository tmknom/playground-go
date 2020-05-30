package main

import (
	"fmt"
)

type report struct {
	sol         int
	temperature temperature
	location    location
}

type location struct {
	lat  float64
	long float64
}

type temperature struct {
	high, low celsius
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type sol int
type report2 struct {
	sol
	temperature
	location
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	return 5
}

func (r report2) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {
	bradbury := location{-4.5859, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Println(report)
	fmt.Println(report.temperature)
	fmt.Println(report.temperature.high)

	fmt.Println(t.average())
	fmt.Println(report.temperature.average())

	fmt.Println("")
	report2 := report2{sol: 15, temperature: t, location: bradbury}
	fmt.Println(report2.temperature)
	fmt.Println(report2.temperature.average())
	fmt.Println(report2.average())
	fmt.Println(report2.high)
	fmt.Println(report2.sol.days(1446))

	fmt.Println("")
	fmt.Println(report2.days(1446))
}
