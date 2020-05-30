package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(temperature)

	temp := temperature["Earth"]
	fmt.Println(temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Println(moon)
	} else {
		fmt.Println("Not found Moon...")
	}

	planets := map[string]string{
		"地球": "Sector ZZ9",
		"火星": "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["地球"] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	delete(planets, "地球")
	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	planets2 := make(map[float64]int, 8)
	fmt.Println(planets2)

	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int)
	for _, t := range temperatures {
		frequency[t]++
	}
	for t, num := range frequency {
		fmt.Printf("%+.2fの出現は%dです\n", t, num)
	}
	fmt.Println(temperatures)
	fmt.Println(frequency)
	fmt.Println("")

	groups := make(map[float64][]float64)
	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}
	for g, temperatures := range groups {
		fmt.Printf("%v : %v\n", g, temperatures)
	}

	fmt.Println("")
	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}
	fmt.Println(set)
	if set[-28.0] {
		fmt.Println("セットのメンバ")
	}

	fmt.Println("")
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	fmt.Println(unique)
	sort.Float64s(unique)
	fmt.Println(unique)
}
