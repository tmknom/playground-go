package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return kelvin(0) // TODO not implement
}

func measure(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Println(k)
	}
}

type sensor func() kelvin

func measure2(samples int, s sensor) {
	measure(samples, s)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	sensor1 := fakeSensor
	fmt.Println(sensor1())

	var sensor2 func() kelvin = realSensor
	fmt.Println(sensor2())

	measure(2, fakeSensor)
	measure2(2, fakeSensor)

	var f = func() {
		fmt.Println("これは無名関数です")
	}
	f()

	func() {
		fmt.Println("無名関数を直接呼び出します")
	}()

	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())

	var k kelvin = 294.0
	clojure := func() kelvin {
		return k
	}
	fmt.Println(clojure())

	k++
	fmt.Println(clojure())
}
