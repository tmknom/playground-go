package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const lightSpeed = 299792
	var distance1, distance2 = 56000000, 401000000
	fmt.Printf("%v 秒\n", distance1/lightSpeed)
	fmt.Printf("%v 秒\n", distance2/lightSpeed)

	fmt.Printf("ランダム1 %v\n", rand.Intn(10)+1)
	fmt.Printf("ランダム2 %v\n", rand.Intn(10)+1)
}
