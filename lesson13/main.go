package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var temperature celsius = 20
	fmt.Println(temperature)

	temperature += 10
	fmt.Println(temperature)

	var warmUp float64 = 10
	temperature += celsius(warmUp)

	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Println(c)
	fmt.Println(k.celsius())

	var f fahrenheit = 294.0
	fmt.Println(f.celsius())
}
